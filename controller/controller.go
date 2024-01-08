package controller

import (
	"fmt"
	"net/http"

	"BlogYmmersion/manager"
	inittemplate "BlogYmmersion/templates"
)

const Port = "localhost:8080"

func TreatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//recupérer les données du formulaire d'enregistrement
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Printf("l'email:%s    le mot de passe:%s \n ", email, password)

		//Vérifier si c'est la première connexion
		if manager.IsFirstLogin(email, password) {
			pseudo := r.FormValue("name")
			//Enregistrer le login
			manager.SaveLogin(pseudo, email, password)
			//Rediriger vers l'acceuil
			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			//Vérifier si les logins ont déjà été enregistrés
			if manager.IsLoginRegistered(email, password) {
				//Rediriger vers l'acceuil
				http.Redirect(w, r, "/home", http.StatusFound)
			} else {
				//Rediriger vers la page d'erreurs
				fmt.Printf("l'email:%s    le mot de passe:%s \n ", email, password)
				http.Redirect(w, r, "/error", http.StatusFound)
			}

		}

	}

}

func ConnexionHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "connexion", nil)
}
func InscriptionHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "inscription", nil)
}
func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "error", nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "home", nil)
}
