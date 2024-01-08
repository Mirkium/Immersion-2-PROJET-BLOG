package routeur

import (
	"BlogYmmersion/controller"
	"log"
	"net/http"
)

func InitServe() {
	fileServe := http.FileServer(http.Dir("assets"))
	http.Handle("assets", http.StripPrefix("/assets/", fileServe))

	http.HandleFunc("/home", controller.HomeHandler)
	http.HandleFunc("/treat", controller.TreatHandler)
	http.HandleFunc("/connexion", controller.ConnexionHandler)
	http.HandleFunc("/connexion", controller.InscriptionHandler)

	if err := http.ListenAndServe(controller.Port, nil); err != nil {

		fmt.printf("ERREUR LORS DE L'INITIATION DES ROUTES %v \n", err)

		log.Fatal(err)

	}
}
