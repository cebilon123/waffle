package packet

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/gopacket/pcap"

	"waffle/internal/worker"
)

var (
	// ErrNetworkInterfaceNotFound is an error which indicates  that network interface was not found
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

// GetNetworkInterface retrieves all available interfaces, verifies if interface's description matches description
// in interfaceDescription field of WindowsNetworkInterfaceProvider struct, and returns an interfaces which
// suits that condition.
func (w *WindowsNetworkInterfaceProvider) GetNetworkInterface() (*pcap.Interface, error) {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		return nil, fmt.Errorf("find all network interfaces, %w", err)
	}

	devicesDescriptions := make([]string, len(interfaces))
	for i, netInterface := range interfaces {
		devicesDescriptions[i] = netInterface.Description

		if netInterface.Description == w.interfaceDescription {
			return &netInterface, nil
		}
	}

	return nil, fmt.Errorf("%w, available devices descriptions: %s", ErrNetworkInterfaceNotFound, strings.Join(devicesDescriptions, " ; "))
}
