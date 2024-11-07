package handlers

import (
	"net/http"
	"webapp/models"
	"webapp/utils"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	utils.RenderTemplate(w, "view", p)
}
