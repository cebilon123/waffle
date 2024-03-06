package guard

import "net/http"

// Defender must be implemented by the struct
// representing defense rule (or set of rules).
type Defender interface {
	Validate(r *http.Request) error
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
