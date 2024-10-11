package view

import (
    "html/template"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
    tpl := `<p>{{ . }}</p>`
    data := "Hello"
    t := template.Must(template.New("a").Parse(tpl))
    t.Execute(w, data)
}
