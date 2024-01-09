package routeur

import (
	"BlogYmmersion/controller"
	"fmt"
	"log"
	"net/http"
)

func InitServe() {
	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))

	http.HandleFunc("/home", controller.HomeHandler)
	http.HandleFunc("/treatmentI", controller.TreatInscriptionHandler)
	http.HandleFunc("/treatmentC", controller.TreatConnexionHandler)
	http.HandleFunc("/connexion", controller.ConnexionHandler)
	http.HandleFunc("/inscription", controller.InscriptionHandler)

	http.HandleFunc("/error", controller.ErrorHandler)
	if err := http.ListenAndServe(controller.Port, nil); err != nil {

		fmt.Printf("ERREUR LORS DE L'INITIATION DES ROUTES %v \n", err)

		log.Fatal(err)

	}
}
