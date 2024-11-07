package main

import (
	"webapp/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", handlers.ViewHandler)
	http.HandleFunc("/edit/", handlers.EditHandler)
	http.HandleFunc("/save/", handlers.SaveHandler)

	http.ListenAndServe(":8080", nil)
}
