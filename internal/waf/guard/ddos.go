package guard

import (
	"context"
	"fmt"

	"waffle/internal/ddos"
	"waffle/internal/request"
)

type DDOS struct {
	ipValidator ddos.IPValidator
}

func (d *DDOS) Validate(ctx context.Context, rw *request.Wrapper) error {
	if err := d.ipValidator.Validate(ctx, rw.IPAddress); err != nil {
		return fmt.Errorf("validate ip using ip validator: %w", err)
	}

	return nil
}
