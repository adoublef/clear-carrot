package jsx

import (
	"net/http"

	g "github.com/maragudk/gomponents"
)

func Render(w http.ResponseWriter, n g.Node) {
	err := n.Render(w)
	if err != nil {
		http.Error(w, "error writing template", http.StatusInternalServerError)
		return
	}
}
