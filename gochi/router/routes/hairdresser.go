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
	db := database.SetupDatabase()
	data, err := database.GetAllHairdressers(db)
	if err != nil {
		SendErrorResponse(w, "Error retrieving hairdressers", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getHairdresserByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	db := database.SetupDatabase()
	data, err := database.GetHairdresserByUID(db, uid)
	if err != nil {
		SendErrorResponse(w, "Hairdresser not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
