package controller

import (
	"BlogYmmersion/manager"
	inittemplate "BlogYmmersion/templates"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const Port = "localhost:8080"

var store = sessions.NewCookieStore([]byte(SecretKey()))

func SecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(key)
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
func TreatInscriptionHandler(w http.ResponseWriter, r *http.Request) {
	//recupérer les données du formulaire d'enregistrement
	email := r.FormValue("email")
	password := r.FormValue("password")

	//Enregistrer le nouvel Utilisateur
	users := manager.RetrieveUser()
	var login bool

	for _, user := range users {
		if user.Email == email && user.Password == password {
			//verifier si le login est déjà enregistré
			login = true
		}
	}
	if login {
		http.Redirect(w, r, "/connexion?error=already_registred", http.StatusFound)
	} else {

		//rediriger vers la page dc'acceuil & enregistrer le login
		manager.MarkLogin(email, password)
		http.Redirect(w, r, "/home?success=Login_registred", http.StatusFound)
	}

}
func TreatConnexionHandler(w http.ResponseWriter, r *http.Request) {
	var session *sessions.Session
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
			break
		}
	}
	if login {
		i := 0
		//Creer une nouvelle session & stocker l'email
		var err error
		session, err = store.Get(r, "session-name")
		for i < 1 {
			if err != nil {
				http.Error(w, "ERREUR DE SESSION_1", http.StatusInternalServerError)
				return
			}
		}

		session.Values["email"] = email
		fmt.Println("EMAIL RECU", email)
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "ERREUR DE SESSION_2", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		//rediriger vers la page de connexion avec un message d'erreur
		http.Redirect(w, r, "/connexion?error=invalid_login_try_again", http.StatusFound)
	}

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
	fmt.Printf("HomeHandler()\tdata: %#v\n", data)

	inittemplate.Temp.ExecuteTemplate(w, "home", data)
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

	//Récuperer l'IP de la catégorie en paramètre de la requetes
	categoryID := r.URL.Query().Get("id")

	//lecture des données json à partir du fichier
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
	if err != nil {
		log.Fatal(err)
		return
	}

	//Trouver la catégorie correspondante à l'ID
	var category manager.Category
	for _, c := range data.Categories {
		if c.ID == categoryID {
			category = c
			break
		}
	}

	inittemplate.Temp.ExecuteTemplate(w, "category", category)
}

// Sécurisation des routes/gestions des erreurs de chargement de pages
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	inittemplate.Temp.ExecuteTemplate(w, "404", nil)
}
func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := manager.LoadComments()
	if err != nil {
		//Gerer l'erreur lors du chargement des commentaires
		log.Println("'erreur lors du chargement des commentaires CommentHandler", err)
		//Rediriger l'utilisateur vers la page d'erreurs
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	//Convertir les commentaires en json
	commentJSON, err := json.Marshal(comments)
	if err != nil {
		log.Println("'erreur lors du Conversion en json ", err)

		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	//en-tete pour indiquer la réponse json
	w.Header().Set("content-Type", "application/json")
	//ECrire les données JSON dans la réponse
	w.Write(commentJSON)
	inittemplate.Temp.ExecuteTemplate(w, "comments", comments)
}

func SubmitCommentHandler(w http.ResponseWriter, r *http.Request) {
	commentaire := r.FormValue("commentaire")
	nomFilm := r.FormValue("nom_film")
	userEmail, err := GetEmailSession(r)
	if err != nil {
		//Rediriger l'utilisateur vers la page de connexion
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}

	comment := manager.Comment{
		Email:       userEmail,
		NomFilm:     nomFilm,
		Commentaire: commentaire,
	}
	comments, err := manager.LoadComments()
	if err != nil {
		//Gerer l'erreur lors du chargement des commentaires
		log.Println("'erreur lors du chargement des commentaires", err)
		//Rediriger l'utilisateur vers la page d'erreurs
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	//Ajout du nouveau commentaire à la liste des commentaires
	comments = append(comments, comment)

	//ENREGISTRER LE COMMENTAIRES mis à jour dans le json
	err = manager.SaveComment(comments)
	if err != nil {
		//Gerer l'erreur lors de la sauvegarde des commentaires
		log.Println("'erreur lors du chargement des commentaires", err)
		//Rediriger l'utilisateur vers la page d'erreurs
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/comments", http.StatusSeeOther)
}

// Récuperer l'email à partir de la session ouverte
func GetEmailSession(r *http.Request) (string, error) {
	session, err := store.Get(r, "session-name")

	if err != nil {
		return "", err
	}
	//Vérifier si le user est authentifié dans la session
	userEmail, ok := session.Values["email"].(string)

	if !ok {
		return "", errors.New("utilisateur non authentifié")
	}
	return userEmail, nil
}
