package routes

import (
	"encoding/json"
	"gohairdresser/database"
	"gohairdresser/structs"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaloonRoutes(r *chi.Mux) {
	r.Route("/saloons", func(r chi.Router) {
		r.Post("/create", createSaloonRoute)
		r.Get("/all", getAllHairSaloons)
		r.Get("/{uid}", getHairSaloonByUID)
	})
}

func getAllHairSaloons(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllHairSaloons()
	if err != nil {
		SendErrorResponse(w, "Error retrieving hair saloons", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getHairSaloonByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairSaloonByUID(uid)
	if err != nil {
		SendErrorResponse(w, "Hair saloon not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}

func createSaloonRoute(w http.ResponseWriter, r *http.Request) {
	var saloon structs.CreateSaloon

	// Decode the body into the struct
	err := json.NewDecoder(r.Body).Decode(&saloon)
	if err != nil {
		if err == io.EOF {
			// Handle empty body
			SendErrorResponse(w, "Request body is empty or in wrong format", err, http.StatusBadRequest)
			return
		}
		// Handle other JSON decoding errors
		SendErrorResponse(w, "Failed to create saloon", err, http.StatusInternalServerError)
		return
	}

	createdSaloon, err := database.CreateSaloon(saloon)
	if err != nil {
		SendErrorResponse(w, "Failed to create saloon", err, http.StatusInternalServerError)
		return
	}

	// Successful response
	SendJSONResponse(w, createdSaloon)
}
