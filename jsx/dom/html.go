package dom

import (
	"net/url"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

// Html struct
type HtmlProps struct {
	Title   string
	BaseURL *url.URL
}

func Html(props HtmlProps, children ...g.Node) g.Node {
	children = append(children, hx.Boost("true"))
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleEl(g.Text(props.Title)),
				// baseUrl
				Link(Rel("preload"), As("script"), Href("/static/htmx.min.js")),
				Script(Src("/static/htmx.min.js"), Defer()),
				Link(Rel("preload"), As("script"), Href("/static/hyperscript.min.js")),
				Script(Src("/static/hyperscript.min.js"), Defer()),
			),
			Body(children...),
		),
	)
}
