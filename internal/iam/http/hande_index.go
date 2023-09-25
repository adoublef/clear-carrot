package http

import (
	"net/http"

	"github.com/adoublef/clear-carrot/jsx"
	"github.com/adoublef/clear-carrot/jsx/dom"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	// hx "github.com/maragudk/gomponents-htmx"
)

func (s *Service) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		props := dom.HtmlProps{
			Title:   "Hello World!",
			BaseURL: r.URL,
		}

		doc := dom.Html(props,
			Header(Nav(
				A(Href("/"), g.Text("home")),
				Ul(
					Li(A(Href("/"), g.Text(("signin")))),
					Li(A(Href("/"), g.Text(("signout")))),
					// Li(A(Href("/settings"),g.Text(("settings")))),
				),
			)),
			Main())

		jsx.Render(w, doc)
	}
}
