package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	mu sync.Mutex
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

func SaveLogin(name string, email string, password string) {
	mu.Lock()
	defer mu.Unlock()

	//Ouvrir le fichier dans lequel iront les logins
	file, err := os.OpenFile("manager/Login.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier:%v", err)
		return
	}
	defer file.Close()

	currentTime := time.Now()

	//
	_, err = fmt.Fprintf(file, "Pseudo:%s, Email: %s, Password: %s, Date: %s\n", name, email, password, currentTime.Format("2006-01-02 15:04:05"))

	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier Login.txt:\n", err)
	}
}

func IsFirstLogin(email string, password string) bool {

	file, err := os.Open("manager/Login.txt")

	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier Login.txt: %v", err)
		return false
	}
	defer file.Close()

	//parcourir le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//Vérifier si la ligne contient les memes informations de login
		if strings.Contains(line, fmt.Sprintf("Email: %s", email)) && strings.Contains(line, fmt.Sprintf("Password: %s", password)) {
			return false // pour indiquer que le login existe déjà/ce n'est pas la première connexion
		}
	}
	//Les infos n'ont pas été trouvées/ c'est la 1ère connexion
	return false
}

func IsLoginRegistered(email string, password string) bool {
	file, err := os.Open("manager/Login.txt")

	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier Login.txt: %v", err)
		return false
	}
	defer file.Close()

	//parcourir le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//Vérifier si la ligne contient les memes informations de login
		if strings.Contains(line, fmt.Sprintf("Email: %s, Password: %s", email, password)) {
			return true
		}
	}
	return false
}
