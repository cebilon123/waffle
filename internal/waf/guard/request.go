package guard

import (
	"net/http"
)

type RequestWrapper struct {
	request *http.Request
}

func NewRequestWrapper(r *http.Request) *RequestWrapper {
	return &RequestWrapper{request: r}
}
