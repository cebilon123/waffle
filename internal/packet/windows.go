package packet

import (
	"errors"
	"fmt"

	"github.com/google/gopacket/pcap"

	"waffle/internal/worker"
)

var (
	ErrNetworkInterfaceNotFound = errors.New("network interface not found")
)

// WindowsNetworkInterfaceProvider is windows network interface provider.
type WindowsNetworkInterfaceProvider struct {
	interfaceDescription string
}

var _ worker.NetworkInterfaceProvider = (*WindowsNetworkInterfaceProvider)(nil)

// NewWindowsNetworkInterfaceProvider creates a WindowsNetworkInterfaceProvider
//
//	interfaceDescription is a identification of the interface i.e.: "WAN Miniport (IP)"
func NewWindowsNetworkInterfaceProvider(interfaceDescription string) *WindowsNetworkInterfaceProvider {
	return &WindowsNetworkInterfaceProvider{interfaceDescription: interfaceDescription}
}

func (w *WindowsNetworkInterfaceProvider) GetNetworkInterface() (*pcap.Interface, error) {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		return nil, fmt.Errorf("find all network interfaces, %w", err)
	}

	for _, netInterface := range interfaces {
		if netInterface.Description == w.interfaceDescription {
			return &netInterface, nil
		}
	}

	return nil, ErrNetworkInterfaceNotFound
}
