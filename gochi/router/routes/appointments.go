package routes

import (
	"github.com/go-chi/chi/v5"
)

func AppointmentsRoutes(r *chi.Mux) {
	r.Route("/appointments", func(r chi.Router) {
	})
}
