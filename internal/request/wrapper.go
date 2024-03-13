package request

import (
	"net/http"
)

type Wrapper struct {
	request *http.Request
}

func NewRequestWrapper(r *http.Request) *Wrapper {
	return &Wrapper{request: r}
}

func (w *Wrapper) Request() *http.Request {
	return w.request
}
