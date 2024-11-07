package handlers

import (
	"net/http"
	"webapp/models"
	"webapp/utils"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title} // Use models.Page instead of Page
	}
	utils.RenderTemplate(w, "edit", p)
}
