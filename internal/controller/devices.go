package controller

import (
	"context"
	"crypto/rand"
	"fmt"
	"net"
	"net/http"
	"net/netip"

	b64 "encoding/base64"

	"entgo.io/ent/dialect/sql"
	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/device"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/ogent"
	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/types"
	"github.com/tiagoposse/connect/internal/utils"
)

// CreateDevice handles POST /devices requests.
func (c *Controller) CreateDevice(ctx context.Context, req *ogent.CreateDeviceReq) (ogent.CreateDeviceRes, error) {

	b := c.client.Device.Create()
	// Add all fields.
	b.SetName(req.Name)
	if v, ok := req.Description.Get(); ok {
		b.SetDescription(v)
	}
	b.SetType(req.Type)

	if len(req.DNS) == 0 {
		b.SetDNS(c.cfg.Wireguard.DnsServers)
	} else {
		b.SetDNS(*ogent.JsonConvert(req.DNS, &types.InetSlice{}).(*types.InetSlice))
	}
	
	b.SetPublicKey(req.PublicKey)
	b.SetKeepAlive(req.KeepAlive)

	// Add all edges.
	b.SetUserID(req.User)
	
	group, err := c.client.Group.Query().Where(group.HasUsersWith(user.IDEQ(req.User))).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting group: %w", err)
	}

	internetCidr := []types.Cidr{{Prefix: netip.MustParsePrefix("0.0.0.0/0")}}
	if len(group.Rules) == 0 {
		b.SetAllowedIps(internetCidr)
	} else {
		allowedIPs := make([]types.Cidr, 0)
		for _, rule := range group.Rules {
			if rule.Type == types.RuleAllow {
				allowedIPs = append(allowedIPs, rule.Target)
			}
		}

		if len(allowedIPs) == 0 {
			b.SetAllowedIps(internetCidr)
		} else {

			b.SetAllowedIps(*ogent.JsonConvert(req.AllowedIps, &types.CidrSlice{}).(*types.CidrSlice))
		}
	}

	if req.Endpoint == "" {
		ds, err := c.client.Device.Query().Select(device.FieldEndpoint).Where(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("%s << $1", device.FieldEndpoint), group.Cidr.String()))
		}).All(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting devices: %w", err)
		}
	
		usedIPs := make([]net.IP, 0)
		for _, dev := range ds {
			usedIPs = append(usedIPs, dev.Endpoint.IP)
		}
	
		next, err := group.Cidr.FindFirstAvailableIP(usedIPs)
		if err != nil {
			return nil, err
		}

		b.SetEndpoint(types.Inet{IP: next})
	} else {
		parsed, err := netip.ParseAddr(req.Endpoint)
		if err != nil {
			return &ogent.R400{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: utils.RawError(err),
			}, err
		}

		if !group.Cidr.Contains(parsed) {
			return &ogent.R400{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: utils.RawError(err),
			}, err
		}

		b.SetEndpoint(types.Inet{}.ParseString(req.Endpoint))
	}

	preSharedKeyBs := make([]byte, 32)
	_, err = rand.Read(preSharedKeyBs)
	if err != nil {
		return nil, fmt.Errorf("generating pre shared key: %w", err)
	}
	preSharedKey := b64.StdEncoding.EncodeToString([]byte(preSharedKeyBs))
	b.SetPresharedKey(preSharedKey)

	edp, _ := b.Mutation().Endpoint()
	if err := c.wg.AddPeer(req.PublicKey, edp.String()); err != nil {
		return nil, fmt.Errorf("creating peer: %w", err)
	}

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: utils.RawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: utils.RawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := c.client.Device.Query().Where(device.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return ogent.NewDeviceCreate(e), nil
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


type Query interface {
	Count(context.Context) (int, error)
}
