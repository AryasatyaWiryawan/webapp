package utils

import (
	"html/template"
	"net/http"
	"webapp/models"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, p)
}
