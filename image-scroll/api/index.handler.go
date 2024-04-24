package api

import (
	"net/http"
	"text/template"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("html/index.html"))
	page.Execute(w, nil)
}

func HTMXHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("library/htmx.min.js"))
	page.Execute(w, nil)
}
