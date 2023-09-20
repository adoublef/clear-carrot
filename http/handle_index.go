package http

import (
	"net/http"

	"github.com/adoublef/clear-carrot/templates"
)

// const (
// 	layoutsDir   = "templates/layouts"
// 	templatesDir = "templates"
// 	extension    = "/*.html"
// )

// var (
// 	//go:embed templates/*
// 	files     embed.FS
// 	templates map[string]*template.Template
// )

// func init() {
// 	if templates == nil {
// 		templates = make(map[string]*template.Template)
// 	}
// 	tmplFiles, err := fs.ReadDir(files, templatesDir)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, tmpl := range tmplFiles {
// 		if tmpl.IsDir() {
// 			continue
// 		}

// 		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name())
// 		if err != nil {
// 			panic(err)
// 		}

// 		templates[tmpl.Name()] = pt
// 	}
// }

func (s *Service) handleIndex() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
            "userAgent": r.UserAgent(),
        }

		templates.IndexPage.Execute(w, data)
	}
}
