package controller

import (
	"context"
	"crypto/rand"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/device"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/ogent"
	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/types"

	"entgo.io/ent/dialect/sql"
)

func (c *Controller) CreateDeviceWire(ctx context.Context, e *ent.DeviceCreate, req *ogent.CreateDeviceReq) error {
	deviceUser, ok := e.Mutation().UserID()
	if !ok {
		return errors.New("device user not provided")
	}
	
	group, err := c.client.Group.Query().Where(group.HasUsersWith(user.IDEQ(deviceUser))).Only(ctx)
	if err != nil {
		return fmt.Errorf("getting group: %w", err)
	}

	ds, err := c.client.Device.Query().Select(device.FieldEndpoint).Where(func(s *sql.Selector) {
		s.Where(sql.ExprP(fmt.Sprintf("%s << $1", device.FieldEndpoint), group.Cidr.String()))
	}).All(ctx)
	if err != nil {
		return fmt.Errorf("getting devices: %w", err)
	}

	usedIPs := make([]net.IP, 0)
	for _, dev := range ds {
		usedIPs = append(usedIPs, dev.Endpoint.IP)
	}
	next, err := group.Cidr.FindFirstAvailableIP(usedIPs)
	if err != nil {
		return err
	}
	
	e.SetEndpoint(types.Inet{IP: next})
	e.SetDNS(c.cfg.Wireguard.DnsServers)

	if len(group.Rules) == 0 {
		e.SetAllowedIps("0.0.0.0/0")
	} else {
		allowedIPs := make([]string, 0)
		for _, rule := range group.Rules {
			if rule.Type == types.RuleAllow {
				allowedIPs = append(allowedIPs, rule.Target)
			}
		}

		if len(allowedIPs) == 0 {
			e.SetAllowedIps("0.0.0.0/0")
		} else {
			e.SetAllowedIps(strings.Join(allowedIPs, ","))
		}
	}

	preSharedKeyBs := make([]byte, 32)
	_, err = rand.Read(preSharedKeyBs)
	if err != nil {
		return fmt.Errorf("generating pre shared key: %w", err)
	}
	preSharedKey := b64.StdEncoding.EncodeToString([]byte(preSharedKeyBs))
	e.SetPresharedKey(preSharedKey)

	pubKey, _ := e.Mutation().PublicKey()
	if err := c.wg.AddPeer(pubKey, next.String()); err != nil {
		return fmt.Errorf("creating peer: %w", err)
	}

	return nil
}

func (c *Controller) DeleteDeviceWire(ctx context.Context, params ogent.DeleteDeviceParams) error {
	dev, err := c.client.Device.Query().Select(device.FieldPublicKey).Where(device.IDEQ(params.ID)).Only(ctx)
	if err != nil {
		return err
	}

	if err := c.wg.RemovePeer(dev.PublicKey); err != nil {
		return fmt.Errorf("creating peer: %w", err)
	}

	return nil
}
