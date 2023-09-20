package http

import (
	"net/http"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	// hx "github.com/maragudk/gomponents-htmx"
)

func (s *Service) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := Page("Hi!", r.URL.Path).Render(w)
		if err != nil {
			http.Error(w, "error parsing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Page(title, currentPath string) g.Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				TitleEl(g.Text(title)),
				Link(Href("/static/index.css"), Rel("stylesheet")),
				Script(Src("/static/htmx.min.js"), Defer()),
				Script(Src("/static/hyperscript.min.js"), Defer()),
			),
			Body(
				Navbar(currentPath),
				H1(g.Text(title)),
				P(g.Textf("Welcome to the page at %v.", currentPath)),
			),
		),
	)
}

func Navbar(currentPath string) g.Node {
	return Nav(hx.Boost("true"),
		NavbarLink("/", "Home", currentPath),
		NavbarLink("/about", "About", currentPath),
	)
}

func NavbarLink(href, name, currentPath string) g.Node {
	return A(Href(href), c.Classes{"is-active": currentPath == href}, g.Text(name))
}
