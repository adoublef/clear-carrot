package http

import (
	"net/http"

	"github.com/adoublef/clear-carrot/jsx/dom"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	// hx "github.com/maragudk/gomponents-htmx"
)

func (s *Service) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := &dom.Html{
			Title:   "Golang 💛 Htmx",
			BaseURL: r.URL,
		}

		err := page.Render(w,
			Header(Nav(
				A(Href("/"), g.Text("home")),
				Ul(
					Li(A(Href("/signin"), g.Text(("signin")))),
					Li(A(Href("/signout"), g.Text(("signout")))),
					// Li(A(Href("/settings"),g.Text(("settings")))),
				),
			)),
			Main())
		if err != nil {
			http.Error(w, "error parsing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
