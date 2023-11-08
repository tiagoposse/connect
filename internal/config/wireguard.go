package config

import (
	"errors"
	"net/netip"

	"github.com/tiagoposse/connect/internal/types"
)

type Wireguard struct {
	Interface        *string         `yaml:"interface"`
	DnsServers       types.InetSlice `yaml:"dnsServers"`
	Cidr             netip.Prefix    `yaml:"cidr"`
	SelfProvisioning bool            `yaml:"selfProvisioning"`
	KeepAlive        bool            `yaml:"keepAlive"`
}

func (wg *Wireguard) Validate() []error {
	errs := make([]error, 0)

	if *wg.Interface == "" {
		errs = append(errs, errors.New("wireguard interface not set"))
	}

	return errs
}
