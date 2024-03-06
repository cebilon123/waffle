package waf

import (
	"log"
	"net/http"
)

type Handler struct {
	next http.Handler
}

func NewHandler(next http.Handler) *Handler {
	return &Handler{next: next}
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("guarding")
	h.next.ServeHTTP(w, r)
}
