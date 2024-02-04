package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gohairdresser/auth"
	"gohairdresser/structs"
)

func AuthentificationRoutes(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		// r.Post("/client_login", clientLogin)
		// r.Post("/saloon_login", saloonLogin)
		r.Post("/client_signup", clientSignup)
		r.Post("/saloon_signup", saloonSignup)
	})
}

// Logins will come later

func clientSignup(w http.ResponseWriter, r *http.Request) {
	var client structs.CreateClient
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		handleJSONDecodingError(w, err)
		return
	}

	uid, err := auth.CreateClient(client)
	if err != nil {
		SendErrorResponse(w, "Error creating client", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, struct {
		UID string `json:"uid"`
	}{
		UID: uid,
	})
}

func saloonSignup(w http.ResponseWriter, r *http.Request) {
	var saloon structs.CreateSaloon
	err := json.NewDecoder(r.Body).Decode(&saloon)
	if err != nil {
		handleJSONDecodingError(w, err)
		return
	}

	uid, err := auth.CreateSaloon(saloon)
	if err != nil {
		SendErrorResponse(w, "Error creating saloon", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, struct {
		UID string `json:"uid"`
	}{
		UID: uid,
	})
}
