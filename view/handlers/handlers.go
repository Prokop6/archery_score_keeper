package handlers

import (
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	index := template.Must(template.ParseFiles("./templates/index.html"))

	index.Execute(w, "")

}

func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	privacy := template.Must(template.ParseFiles("./templates/privacy.html"))

	privacy.Execute(w, "")

}
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	contact := template.Must(template.ParseFiles("./templates/contact.html"))

	contact.Execute(w, "")

}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	about := template.Must(template.ParseFiles("./templates/about.html"))

	about.Execute(w, "")

}

func NewSessionHandler(w http.ResponseWriter, r *http.Request) {

	session_settings := template.Must(template.ParseFiles("./templates/session_settings.html"))

	session_settings.Execute(w, "")
}
