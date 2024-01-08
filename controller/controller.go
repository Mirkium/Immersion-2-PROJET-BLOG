package controller

import (
	"net/http"

	inittemplate "BlogYmmersion/templates"
)

const Port = "localhost:8080"

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "home", nil)
}
func ConnexionHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}
func InscriptionHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}
