package ddosml

import (
	"context"
	"fmt"
	"net/http"
)

type RequestRepository interface {
	CreateRequest(ctx context.Context, req *Request) error
}

type Request struct {
	IsDDOS bool
}

// DDOS represents machine learning
// DDOS protection.
type DDOS struct {
	// isEnabled is used in order to enable or disable the ddosml
	isEnabled  bool
	repository RequestRepository
}

// NewDDOS creates new ddos ML analyzer used to
// analyze requests in order to find out if given
// request is ddos attack or not.
func NewDDOS(isEnabled bool, repository RequestRepository) *DDOS {
	return &DDOS{
		isEnabled:  isEnabled,
		repository: repository,
	}
}

// IsRequestSuspicious checks if given request is suspicious (and then saves it in the database in order to be
// used in future evaluations of this validator)
func (d *DDOS) IsRequestSuspicious(ctx context.Context, req *http.Request) (bool, error) {
	ok, err := d.validateRequest(ctx, req)
	if err != nil {
		return ok, fmt.Errorf("validate request: %w", err)
	}

	return ok, nil
}

// validateRequest is used to validate if given request is ddos or not. The validation is based on the ML
// model, which decides based on normal user decisions from the UI.
func (d *DDOS) validateRequest(ctx context.Context, req *http.Request) (bool, error) {

}
