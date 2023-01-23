package web

import (
	"html/template"
	"net/http"
)

func ShowLandingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/welcome.html"))
	tmpl.Execute(w, nil)
}
