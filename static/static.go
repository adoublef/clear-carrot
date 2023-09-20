package static

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed *.js *.css
	files embed.FS
	fsys  fs.FS
)

func init() {
	var err error
	if fsys, err = fs.Sub(files, "."); err != nil {
		panic(err)
	}
}

type Static struct {
	Prefix string
}

func (s *Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix(s.Prefix, http.FileServer(http.FS(fsys))).ServeHTTP(w, r)
}
