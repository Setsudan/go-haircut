package routes

import (
	"gohairdresser/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaloonRoutes(r *chi.Mux) {
	r.Route("/saloons", func(r chi.Router) {
		r.Get("/all", getAllHairSaloons)
		r.Get("/{uid}", getHairSaloonByUID)
	})
}

func getAllHairSaloons(w http.ResponseWriter, r *http.Request) {
	db := database.SetupDatabase()
	data, err := database.GetAllHairSaloons(db)
	if err != nil {
		SendErrorResponse(w, "Error retrieving hair saloons", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getHairSaloonByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	db := database.SetupDatabase()
	data, err := database.GetHairSaloonByUID(db, uid)
	if err != nil {
		SendErrorResponse(w, "Hair saloon not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
