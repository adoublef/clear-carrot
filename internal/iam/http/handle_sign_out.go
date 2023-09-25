package http

import "net/http"

func (s *Service) handleSignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
