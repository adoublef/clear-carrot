package templates

import (
	"embed"
	"html/template"
)

var (
	//go:embed index.html
	indexFS embed.FS
)

var IndexPage *template.Template

func init() {
	IndexPage = template.Must(template.ParseFS(indexFS, "index.html"))
}

/* 

var (
	Index iota
)

type Templates struct {}

*/

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

// func handle(responseWriter, data, code)