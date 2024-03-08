package waf

import (
	"log"
	"net/http"

	"waffle/internal/waf/guard"
)

type Handler struct {
	next     http.Handler
	defender guard.Defender
}

func NewHandler(
	next http.Handler,
	defender guard.Defender,
) *Handler {
	return &Handler{
		next:     next,
		defender: defender,
	}
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("guarding")

	if err := h.defender.Validate(guard.NewRequestWrapper(r)); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	h.next.ServeHTTP(w, r)
}
