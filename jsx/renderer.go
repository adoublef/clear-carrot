package node

import (
	"io"

	g "github.com/maragudk/gomponents"
)


type Renderer interface {
	Render (w io.Writer, children ...g.Node) error
}