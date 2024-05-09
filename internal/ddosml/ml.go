package ddosml

import (
	"context"
	"fmt"
	"net/http"
)

type Validator interface {
	ValidateRequest(ctx context.Context, req *http.Request) (bool, error)
}

// DDOS represents machine learning
// DDOS protection.
type DDOS struct {
	// isEnabled is used in order to enable or disable the ddosml
	isEnabled bool
	validator Validator
}

// NewDDOS creates new ddos ML analyzer used to
// analyze requests in order to find out if given
// request is ddos attack or not.
func NewDDOS(isEnabled bool, validator Validator) *DDOS {
	return &DDOS{
		isEnabled: isEnabled,
		validator: validator,
	}
}

// IsRequestSuspicious checks if given request is suspicious (and then saves it in the database in order to be
// used in future evaluations of this validator)
func (d *DDOS) IsRequestSuspicious(ctx context.Context, req *http.Request) (bool, error) {
	ok, err := d.validator.ValidateRequest(ctx, req)
	if err != nil {
		return ok, fmt.Errorf("validator validate request: %w", err)
	}

	return ok, nil
}
