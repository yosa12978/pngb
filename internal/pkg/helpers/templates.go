package helpers

import (
	"html/template"
	"io"
)

func RenderTemplate(w io.Writer, templ string, payload interface{}) {
	t, _ := template.ParseFiles(templ)
	t.Execute(w, payload)
}
