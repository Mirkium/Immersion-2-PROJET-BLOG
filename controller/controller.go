package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"BlogYmmersion/manager"
	inittemplate "BlogYmmersion/templates"
)

const Port = "localhost:8080"

func TreatInscriptionHandler(w http.ResponseWriter, r *http.Request) {
	//recupérer les données du formulaire d'enregistrement
	email := r.FormValue("email")
	password := r.FormValue("password")

	//Enregistrer le nouvel Utilisateur
	manager.MarkLogin(email, password)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func TreatConnexionHandler(w http.ResponseWriter, r *http.Request) {
	//recupérer les données du formulaire de connexion
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println("l' email:", email)
	fmt.Println("le password:", password)
	users := manager.RetrieveUser()
	var login bool

	for _, user := range users {
		if user.Email == email && user.Password == password {
			//verifier si le login est correcte
			login = true

		}
	}
	if login {
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {

		//rediriger vers la page de connexion avec un message d'erreur
		http.Redirect(w, r, "/connexion?error=invalid_login", http.StatusFound)
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

	//ouverture et lecture des données json à partir du fichier
	dataJSON, err := os.ReadFile("DATA.json")

	//gestion de l'erreur à la lecture du fichier
	if err != nil {
		log.Fatal(err)
		manager.PrintColorResult("red", " ERREUR LORS DE LA LECTURE DU FICHIER")
		fmt.Printf("ALERTE: %#v", err)
		return
	}

	//Désérialiser les données JSON dans
	// la structure de données(Analyse des reponses json)
	var data manager.DataCategory
	err = json.Unmarshal(dataJSON, &data)

	//gestion de l'erreur à la lecture du fichier
	if err != nil {
		log.Fatal(err)
		manager.PrintColorResult("red", " ERREUR LORS DE LA LECTURE DU FICHIER")
		fmt.Printf("ALERTE: %#v", err)
		return
	}

	inittemplate.Temp.ExecuteTemplate(w, "home", data)
}
