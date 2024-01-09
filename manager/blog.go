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

// structure de sauvegarde du login de chaque  user
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ListUser []LoginUser
)

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

// func SaveLogin() error {

// 	//Ouvrir le fichier dans lequel iront les logins
// 	file, err := os.OpenFile("manager/Login.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

// 	if err != nil {
// 		fmt.Printf("Erreur lors de la lecture du fichier:%v", err)
// 		return err
// 	}

// Verifier si le user est déjà enregistré ou pas
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
	//
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("manager/Login.txt", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("list des users : %#v\n", users)

}
