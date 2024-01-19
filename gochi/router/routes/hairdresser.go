package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HairdresserRoutes(r *chi.Mux) {
	r.Route("/hairdressers", func(r chi.Router) {
		r.Put("/{hairdresserId}/schedule", updateHairdresserSchedule) // as defined previously
	})
}

func updateHairdresserSchedule(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHairdresserSchedule")
}
