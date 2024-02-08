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
		r.Delete("/{uid}", deleteSaloonRoute)
	})
}

func getAllHairSaloons(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllHairSaloons()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving hair saloons", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hair saloons retrieved successfully", data, nil)
}

func getHairSaloonByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairSaloonByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Hair saloon not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hair saloon retrieved successfully", data, nil)
}

func createSaloonRoute(w http.ResponseWriter, r *http.Request) {
	var saloon structs.CreateSaloon

	err := json.NewDecoder(r.Body).Decode(&saloon)
	if err != nil {
		message := "Failed to decode request body"
		statusCode := http.StatusInternalServerError
		if err == io.EOF {
			message = "Request body is empty or in wrong format"
			statusCode = http.StatusBadRequest
		}
		SendResponse(w, statusCode, "Error", message, nil, err)
		return
	}

	createdSaloon, err := database.CreateSaloon(saloon)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Failed to create saloon", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon created successfully", createdSaloon, nil)
}

func deleteSaloonRoute(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	err := database.DeleteSaloon(uid)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Failed to delete saloon", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon deleted successfully", nil, nil)
}
