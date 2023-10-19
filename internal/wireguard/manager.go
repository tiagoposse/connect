package wireguard

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/tiagoposse/connect/internal/config"
)

type WireGuardManager struct {
	InterfaceName string
	mu            sync.Mutex
	cidr          string
	dnsServers    []string
	changes int
	tmp string
}

func NewWireGuardManager(cfg *config.Wireguard) *WireGuardManager {
	p, err := os.MkdirTemp("/tmp", "wg")
	if err != nil {
		panic(err)
	}

	return &WireGuardManager{
		InterfaceName: cfg.Interface,
		cidr:          cfg.Cidr,
		dnsServers:    cfg.DnsServers,
		changes: 0,
		tmp: p,
	}
}

func (m *WireGuardManager) AddPeer(pubKey, endpoint string) error {
	m.mu.Lock()
	defer m.mu.Unlock()


	cmd := exec.Command("wg", "set", m.InterfaceName, "peer", pubKey, "allowed-ips", fmt.Sprintf("%s/32", endpoint))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add peer: %w", err)
	}

	fmt.Println("added the peer")
	m.changes += 1

	return m.reloadInterface()
}

func (m *WireGuardManager) RemovePeer(publicKey string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	cmd := exec.Command("wg", "set", m.InterfaceName, "peer", publicKey, "remove")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove peer: %w", err)
	}
	m.changes += 1

	return m.reloadInterface()
}

func (m *WireGuardManager) reloadInterface() error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("wg syncconf %s <(wg-quick strip %s)", m.InterfaceName, m.InterfaceName))
	if out, err := cmd.Output(); err != nil {
		return fmt.Errorf("failed to reload interface: %s", string(out))
	}

	fmt.Println("done")
	m.changes = 0

	return nil
}
