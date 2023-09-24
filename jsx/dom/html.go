package dom

import (
	"io"
	"net/url"

	"github.com/adoublef/clear-carrot/jsx"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

var _ jsx.Renderer = (*Html)(nil)

// Html struct
type Html struct {
	Title string
	BaseURL *url.URL
}

func (h *Html) Render(w io.Writer, children ...g.Node) error {
	return htmx(h.Title, h.BaseURL, children...).Render(w)
}

func htmx(title string, baseUrl *url.URL, children ...g.Node) g.Node {
	children = append(children, hx.Boost("true"))
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleEl(g.Text(title)),
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
