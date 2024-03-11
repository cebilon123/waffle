package waf

import (
	"fmt"
	"net/http"
	"time"

	"waffle/internal/ratelimit"
	"waffle/internal/request"
	"waffle/internal/waf/guard"
)

type Handler struct {
	next     http.Handler
	defender guard.Defender
	limiter  ratelimit.Limiter
}

func NewHandler(
	next http.Handler,
	defender guard.Defender,
	limiter ratelimit.Limiter,
) *Handler {
	return &Handler{
		next:     next,
		defender: defender,
		limiter:  limiter,
	}
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ipAddr, err := request.GetRealIPAddress(*r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(fmt.Errorf("resolve address for request: %w", err).Error()))
		return
	}

	rate := h.limiter.GetRate(r.Context(), ipAddr)
	if rate != nil && rate.IsLimited() {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	tmp := h.limiter.SetRate(r.Context(), ipAddr, time.Now().Add(time.Second*5))
	_, _ = w.Write([]byte(tmp))

	if err := h.defender.Validate(guard.NewRequestWrapper(r)); err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	h.next.ServeHTTP(w, r)
}
