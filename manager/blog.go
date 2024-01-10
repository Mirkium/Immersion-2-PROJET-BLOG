package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// structure globale de chaque catégorie
type Category struct {
	Image string `json:"image"`
	ID    string `json:"id"`
	Name  string `json:"name"`

	Films []Film
}

// structure globale de l'objet film
type Film struct {
	ID          string `json:"id"`
	Titre       string `json:"titre"`
	Auteur      string `json:"auteur"`
	Synopsis    string `json:"synopsis"`
	Image       string `json:"image"`
	Note        string `json:"note"`
	RealeseDate string `json:"realese_date"`
}
type DataCategory struct {
	Categories []Category
}

type Comment struct {
	Email       string `json:"email"`
	NomFilm     string `json:"nom_film"`
	Commentaire string `json:"Commentaire"`
}

// structure de sauvegarde du login de chaque  user
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ListUser []LoginUser

const CommentFile = "manager/comments.txt"

func PrintColorResult(color string, message string) {
	colorCode := ""
	switch color {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	case "yellow":
		colorCode = "\033[33m"
	case "blue":
		colorCode = "\033[34m"
	case "purple":
		colorCode = "\033[35m"

	default: //REMETTRE LA COULEUR INITIALE (blanc)
		colorCode = "\033[0m"
	}
	fmt.Printf("%s%s\033[0m", colorCode, message)
}

func RetrieveUser() []LoginUser {
	data, err := os.ReadFile("manager/Login.txt")

	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier:%v", err)
		return nil
	}
	var Users []LoginUser
	err = json.Unmarshal(data, &Users)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("list des users : %#v\n", Users)
	return Users
}

// Marquer ( enregistrer) les nouveaux users dans le fichiers De login
func MarkLogin(email string, password string) {
	var newLogin = LoginUser{
		Email:    email,
		Password: password,
	}
	users := RetrieveUser()
	users = append(users, newLogin)

	//Convertir lelogin en JSON
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	//Ecrire les données JSON dans le fichier
	err = os.WriteFile("manager/Login.txt", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("list des users : %#v\n", users)

}

// Enregistrer les commentaires
func SaveComment(newComment []Comment) error {

	//Charger les commentaires déjà enregistrés
	comments, err := LoadComments()
	if err != nil {
		return err
	}
	//Ajouter un nouveau commentaire
	comments = append(comments, newComment...)

	//Convertir lelogin en JSON
	data, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	//Ecrire les données JSON dans le fichier
	err = os.WriteFile(CommentFile, data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("liste des commentaires : %#v\n", newComment)
	return nil
}

// Charger les commentaires à partir d'un fichier json
func LoadComments() ([]Comment, error) {
	//Vérifier si le fichier json

	_, err := os.Stat(CommentFile)
	if os.IsNotExist(err) {
		return []Comment{}, nil
	} else if err != nil {
		fmt.Printf("Erreur lors de la verification du fichier : %#v\n", err)
		return nil, err
	}

	//lecture des données du fichier Json
	dataJSON, err := os.ReadFile(CommentFile)
	if err != nil {
		log.Fatal(err)
	}
	//désérialiser les données json en une liste de commentaires
	var comments []Comment
	err = json.Unmarshal(dataJSON, &comments)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("list des commentaires : %#v\n", comments)
	return comments, err
}
