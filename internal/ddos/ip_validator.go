package ddos

import (
	"context"
	"errors"
	"net"
	"time"
)

var (
	fetchInterval = time.Hour * 12
)

// IPValidator should be implemented by the structs
// that are able to validate ip addresses in some way.
type IPValidator interface {
	// Validate returns error if ip is invalid / malicious etc.
	Validate(ctx context.Context, ip *net.IP) error
}

type TreeSetIPValidator struct {
	setProvider IPTreeSetProvider
}

func NewTreeSetIPValidator() *TreeSetIPValidator {
	return &TreeSetIPValidator{
		setProvider: NewSyncIPTreeSetProvider(fetchInterval, fetchFunc),
	}
}

func (t *TreeSetIPValidator) Validate(_ context.Context, ip *net.IP) error {
	set := t.setProvider.GetSet()
	if set == nil {
		return errors.New("ip addresses tree set is nil")
	}

	if set.Contains(*ip) {
		return errors.New("given IP is in the malicious collection")
	}

	return nil
}

func fetchFunc(ctx context.Context) ([]net.IP, error) {
	return nil, nil
}
