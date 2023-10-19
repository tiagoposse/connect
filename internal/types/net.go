package types

import (
	"database/sql/driver"
	"fmt"
	"net"
	"net/netip"
)

// Inet represents a single IP address
type Inet struct {
	net.IP
}

// Scan implements the Scanner interface
func (i *Inet) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
			if i.IP = net.ParseIP(string(v)); i.IP == nil {
					err = fmt.Errorf("invalid value for ip %q", v)
			}
	case string:
			if i.IP = net.ParseIP(v); i.IP == nil {
					err = fmt.Errorf("invalid value for ip %q", v)
			}
	default:
			err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface
func (i Inet) Value() (driver.Value, error) {
	return i.IP.String(), nil
}

func (i Inet) String() string {
	return i.IP.String()
}

func (i Inet) ParseString(s string) Inet {
	err := i.Scan(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Cidr represents a Cidr range
type Cidr struct {
	netip.Prefix
}

// Scan implements the Scanner interface
func (c *Cidr) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		c.Prefix, err = netip.ParsePrefix(string(v))
	case string:
		c.Prefix, err = netip.ParsePrefix(v)
	default:
			err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface
func (i Cidr) Value() (driver.Value, error) {
	return i.Prefix.String(), nil
}

func (i Cidr) String() string {
	return i.Prefix.String()
}

func (i Cidr) ParseString(s string) Cidr {
	err := i.Scan(s)
	if err != nil {
		panic(err)
	}
	return i
}


func (i Cidr) FindFirstAvailableIP(usedIPs []net.IP) (net.IP, error) {
	_, ipNet, err := net.ParseCIDR(i.Prefix.String())
	if err != nil {
		return nil, err
	}

	// Convert the network IP to a 4-byte representation
	ip := ipNet.IP.To4()

	// Create a map for faster lookup of used IPs
	used := make(map[string]bool)
	for _, usedIP := range usedIPs {
		used[usedIP.String()] = true
	}

	// Iterate through the IPs in the CIDR range
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		// Skip the network address and broadcast address
		if ip[3] == 0 || ip[3] == 255 {
			continue
		}

		// Check if this IP is used
		if !used[ip.String()] {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("no available IPs in the given CIDR range")
}

// Increment an IPv4 address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
