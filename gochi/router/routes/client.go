package routes

import (
	"github.com/go-chi/chi/v5"
)

func ClientsRoutes(r *chi.Mux) {
	r.Route("/clients", func(r chi.Router) {
	})
}
