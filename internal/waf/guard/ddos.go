package guard

import "waffle/internal/ddos"

type DDOS struct {
	ipValidator ddos.IPValidator
}

func (D *DDOS) Validate(rw *RequestWrapper) error {
	//TODO implement me
	panic("implement me")
}
