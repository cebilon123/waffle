package guard

import (
	"fmt"
	"sync"

	"waffle/internal/request"
)

// Defender must be implemented by the struct
// representing defense rule (or set of rules).
type Defender interface {
	Validate(rw *request.Wrapper) error
}

// DefenseCoordinator coordinates defense. It validates request against set of defenders.
// it also implements Defender interface, but the struct is treated like the commander
// (that defends against attacks).
type DefenseCoordinator struct {
	defenders []Defender
}

func NewDefenseCoordinator(defenders []Defender) *DefenseCoordinator {
	return &DefenseCoordinator{defenders: defenders}
}

func (d *DefenseCoordinator) Validate(rw *request.Wrapper) error {
	var wg sync.WaitGroup

	errChan := make(chan error)

	go func() {
		defer close(errChan)

		for _, defender := range d.defenders {
			wg.Add(1)

			go func(rw *request.Wrapper, d Defender, errChan chan error) {
				if err := d.Validate(rw); err != nil {
					errChan <- err
				}

				wg.Done()
			}(rw, defender, errChan)
		}

		wg.Wait()
	}()

	select {
	case <-rw.Request().Context().Done():
		return nil
	case err, ok := <-errChan:
		if !ok {
			return nil
		}

		return fmt.Errorf("defender: %w", err)
	}
}
