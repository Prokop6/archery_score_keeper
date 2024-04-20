package handlers

import (
	"html/template"
	"net/http"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	static := template.Must(template.ParseFiles("./static/style.css"))

	w.Header().Add("Content-Type", "text/css")
	static.Execute(w, "")

}