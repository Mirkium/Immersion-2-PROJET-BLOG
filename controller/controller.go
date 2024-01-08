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

	inittemplate.Temp.ExecuteTemplate(w, "connexion", nil)
}
func InscriptionHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "inscription", nil)
}

func TreatHandler(w http.ResponseWriter, r *http.Request) {

	// http.Redirect(w http.ResponseWriter , r *http.Request, "/acceuil",ht)
}
