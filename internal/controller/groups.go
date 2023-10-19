package controller

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/ogent"
	"github.com/tiagoposse/connect/internal/types"
	"github.com/tiagoposse/connect/internal/utils"
)

func (c *Controller) CreateGroupWire(ctx context.Context, g *ent.GroupCreate, req *ogent.CreateGroupReq) error {
	cidr, err := netip.ParsePrefix(req.Cidr)
	if err != nil {
		return err
	}

	rules := make([]utils.IpTablesOperation, 0)
	for _, r := range req.Rules {
		targetPref, err := netip.ParsePrefix(r.Target)
		if err != nil {
			return err
		}
		rules = append(rules, utils.IpTablesOperation{
			Action: types.RuleType(r.Type),
			Destination: types.Cidr{Prefix: targetPref},
			Origin: types.Cidr{Prefix: cidr},
			Operation: utils.IpTablesAdd,
		})
	}

	return utils.ExecRulesWithRollback(rules)
}

func (c *Controller) UpdateGroupWire(ctx context.Context, e *ent.GroupUpdateOne, req *ogent.UpdateGroupReq, params ogent.UpdateGroupParams) error {
	if req.Rules == nil {
		return nil
	}

	existing, err := c.client.Group.Get(ctx, params.ID)
	if err != nil {
		return err
	}
	
	rules := make([]utils.IpTablesOperation, 0)
	for _, er := range existing.Rules {
		found := false
		for _, nr := range req.Rules {
			pref, err := netip.ParsePrefix(nr.Target)
			if err != nil {
				return err
			}

			if er.Target.Prefix == pref && er.Type == types.RuleType(nr.Type) {
				found = true
				break
			}
		}
		if !found {
			rules = append(rules, utils.IpTablesOperation{
				Operation: utils.IpTablesDrop,
				Origin: existing.Cidr,
				Destination: er.Target,
				Action: er.Type,
			})
		}
	}

	for _, nr := range req.Rules {
		for _, er := range existing.Rules {
			found := false
			pref, err := netip.ParsePrefix(nr.Target)
			if err != nil {
				return err
			}
			if er.Target.Prefix == pref && er.Type == types.RuleType(nr.Type) {
				found = true
				break
			}

			var origin types.Cidr
			if val, ok := req.Cidr.Get(); ok {
				newCidr, err := netip.ParsePrefix(val)
				if err != nil {
					return fmt.Errorf("new cidr %s is not valid: %w", val, err)
				}
				origin = types.Cidr{Prefix: newCidr}
			} else {
				origin = existing.Cidr
			}

			if !found {
				rules = append(rules, utils.IpTablesOperation{
					Operation: utils.IpTablesAdd,
					Origin: origin,
					Destination: types.Cidr{Prefix: pref},
					Action: types.RuleType(nr.Type),
				})
			}
		}
	}

	return utils.ExecRulesWithRollback(rules)
}

func (c *Controller) DeleteGroupWire(ctx context.Context, params ogent.DeleteGroupParams) error {
	existing, err := c.client.Group.Get(ctx, params.ID)
	if err != nil {
		return err
	}
	
	rules := make([]utils.IpTablesOperation, 0)
	for _, r := range existing.Rules {
		rules = append(rules, utils.IpTablesOperation{
			Action: types.RuleType(r.Type),
			Destination: r.Target,
			Origin: existing.Cidr,
			Operation: utils.IpTablesAdd,
		})
	}
	return utils.ExecRulesWithRollback(rules)
}
