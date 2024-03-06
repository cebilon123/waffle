package guard

import (
	"net/http"
	"sync"
)

type RequestWrapper struct {
	request *http.Request

	mu sync.Mutex
}

func NewRequestWrapper(r *http.Request) *RequestWrapper {
	return &RequestWrapper{request: r}
}
