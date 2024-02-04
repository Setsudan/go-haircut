package routes

import (
	"gohairdresser/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HairdresserRoutes(r *chi.Mux) {
	r.Route("/hairdressers", func(r chi.Router) {
		r.Get("/all", getAllHairdressers)
		r.Get("/{uid}", getHairdresserByUID)
	})
}

func getAllHairdressers(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllHairdressers()
	if err != nil && data == nil {
		SendErrorResponse(w, "Error retrieving hairdressers", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getHairdresserByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairdresserByUID(uid)
	if err != nil {
		SendErrorResponse(w, "Hairdresser not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
