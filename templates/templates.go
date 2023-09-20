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