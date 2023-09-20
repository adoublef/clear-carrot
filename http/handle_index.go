package http

import (
	"net/http"

	"github.com/adoublef/clear-carrot/templates"
)

func (s *Service) handleIndex() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
            "userAgent": r.UserAgent(),
        }

		templates.IndexPage.Execute(w, data)
	}
}
