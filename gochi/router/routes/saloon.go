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

func createSaloonRoute(w http.ResponseWriter, r *http.Request) {
	var saloon structs.CreateHairSaloon

	// Decode the body into the struct
	err := json.NewDecoder(r.Body).Decode(&saloon)
	if err != nil {
		if err == io.EOF {
			// Handle empty body
			response := structs.APIResponse{
				Code:    http.StatusBadRequest,
				Status:  "error",
				Message: "Request body is empty or in wrong format",
				Data:    structs.CreateHairSaloon{}, // Provide an example of the expected format
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		// Handle other JSON decoding errors
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the saloon into the database
	createdSaloon, err := database.CreateSaloon(saloon)
	if err != nil {
		response := structs.APIResponse{
			Code:    http.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to create saloon",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Successful response
	response := structs.APIResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "Saloon created successfully",
		Data:    createdSaloon,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
