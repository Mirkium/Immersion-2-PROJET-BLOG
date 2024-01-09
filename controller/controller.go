package controller

import (
	"net/http"

	"BlogYmmersion/manager"
	inittemplate "BlogYmmersion/templates"
)

const Port = "localhost:8080"

func TreatInscriptionHandler(w http.ResponseWriter, r *http.Request) {
	//recupérer les données du formulaire d'enregistrement
	email := r.FormValue("email")
	password := r.FormValue("password")
	pseudo := r.FormValue("name")

	//Appeler la fonction SaveLogin pour Enregistrer le login
	err := manager.SaveLogin(pseudo, email, password)
	if err != nil {
		// Gérer l'erreur, par exemple, afficher un message d'erreur à l'utilisateur
		http.Error(w, "Une erreur s'est produite lors de la sauvegarde des informations de connexion.", http.StatusInternalServerError)
		return
	}
	//rediriger vers l'acceuil
	http.Redirect(w, r, "/home", http.StatusFound)
}
func TreatConnexionHandler(w http.ResponseWriter, r *http.Request) {
	//recupérer les données du formulaire de connexion
	email := r.FormValue("email")
	password := r.FormValue("password")

	//verifier si le login existe dans fichier
	found := manager.CheckLogin(email, password)
	if !found {
		//rediriger vers la page de connexion avec un message d'erreur
		http.Redirect(w, r, "/connexion?error=invalid_login", http.StatusFound)
		return
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
