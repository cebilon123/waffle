package handler

import (
	"net/http"

	"waffle/internal/domain"
)

// RedirectHandler is responsible for redirecting clients to desired servers.
func RedirectHandler(dns domain.NameSystemProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
