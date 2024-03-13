package guard

import (
	"waffle/internal/ddos"
	"waffle/internal/request"
)

type DDOS struct {
	ipValidator ddos.IPValidator
}

func (D *DDOS) Validate(rw *request.Wrapper) error {
	//TODO implement me
	panic("implement me")
}
