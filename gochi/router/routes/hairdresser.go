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
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving hairdressers", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdressers retrieved successfully", data, nil)
}

func getHairdresserByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairdresserByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Hairdresser not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdresser retrieved successfully", data, nil)
}
