package ddos

import (
	"bytes"
	"context"
	"errors"
	"net"

	"github.com/emirpasic/gods/sets/treeset"
)

// IPValidator should be implemented by the structs
// that are able to validate ip addresses in some way.
type IPValidator interface {
	// Validate returns error if ip is invalid / malicious etc.
	Validate(ctx context.Context, ip *net.IP) error
}

type TreeSetIPValidator struct {
	set *treeset.Set
}

func NewTreeSetIPValidator(ipAddresses []net.IP) *TreeSetIPValidator {
	return &TreeSetIPValidator{
		set: treeset.NewWith(ipAddressComparator, ipAddresses),
	}
}

func (t *TreeSetIPValidator) Validate(_ context.Context, ip *net.IP) error {
	if t.set.Contains(*ip) {
		return errors.New("given IP is in the malicious collection")
	}

	return nil
}

func ipAddressComparator(a, b interface{}) int {
	aIP := a.(net.IP)
	bIP := b.(net.IP)

	return bytes.Compare(aIP, bIP)
}
