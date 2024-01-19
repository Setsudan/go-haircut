package routes

import (
	"github.com/go-chi/chi/v5"
)

func AuthentificationRoutes(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
	})
}
