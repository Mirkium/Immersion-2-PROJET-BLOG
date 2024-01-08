package controller

import (
	"net/http"

	inittemplate "BlogYmmersion/templates"
)

const Port = "localhost:8080"

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}
func AdminHandler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}
func Handler(w http.ResponseWriter, r *http.Request) {

	inittemplate.Temp.ExecuteTemplate(w, "", nil)
}
