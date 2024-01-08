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
	http.HandleFunc("/result", controller.AdminHandler)
	http.HandleFunc("/result", controller.ConnexionHandler)
	http.HandleFunc("/result", controller.InscriptionHandler)
	if err := http.ListenAndServe(controller.Port, nil); err != nil {
		log.Fatal(err)
	}
}
