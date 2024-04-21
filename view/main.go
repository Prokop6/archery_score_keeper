package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Prokop6/archery_score_keeper/view/handlers"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.RootHandler)
	r.Get("/style.css", handlers.StyleHandler)
	r.Get("/about", handlers.AboutHandler)
	r.Get("/privacy", handlers.PrivacyHandler)
	r.Get("/contact", handlers.ContactHandler)

	r.Get("/session/settings", handlers.NewSessionHandler)

	log.Print("Server online!")
	http.ListenAndServe(":8080", r)
	log.Print("Server offline...")
}
