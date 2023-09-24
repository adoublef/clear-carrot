package dom

import (
	"io"
	"net/url"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/html"
)

type Renderer interface {
	Render (w io.Writer, children ...g.Node) error
}

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
	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				// baseUrl
				html.TitleEl(g.Text(title)),
				// link script
				html.Script(html.Src("/static/htmx.min.js"), html.Defer()),
				// link script
				html.Script(html.Src("/static/hyperscript.min.js"), html.Defer()),
			),
			html.Body(children...),
		),
	)
}
