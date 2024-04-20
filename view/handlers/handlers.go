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
	static := template.Must(template.ParseFiles("./templates/privacy.html"))

	static.Execute(w, "")

}
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	static := template.Must(template.ParseFiles("./templates/contact.html"))

	static.Execute(w, "")

}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	static := template.Must(template.ParseFiles("./templates/about.html"))

	static.Execute(w, "")

}
