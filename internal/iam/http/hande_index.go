package http

import (
	"fmt"
	"net/http"
	"strconv"

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
			Body(hx.Boost("true"),
				Navbar(currentPath),
				H1(g.Text(title)),
				P(g.Textf("Welcome to the page at %v.", currentPath)),
				Counter(0),
			),
		),
	)
}

func Navbar(currentPath string) g.Node {
	anchor := func(href, name, currentPath string) g.Node {
		return A(Href(href), c.Classes{"is-active": currentPath == href}, g.Text(name))
	}

	return Nav( anchor("/", "Home", currentPath), anchor("/about", "About", currentPath))
}

func (s *Service) handleCount() http.HandlerFunc {
	parseInt := func(r *http.Request) int {
		count := r.URL.Query().Get("count")
		if count == "" {
			count = "0"
		}
		n, _ := strconv.Atoi(count)

		return n
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := Counter(parseInt(r)).Render(w)
		if err != nil {
			http.Error(w, "error parsing counter template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Counter(count int) g.Node {
	url := fmt.Sprintf("/?count=%d", count+1)
	return Span(g.Textf("%d", count),
		Button(g.Text("click"), hx.Post(url), hx.Target("closest span"), hx.Swap("outerHTML")))
}
