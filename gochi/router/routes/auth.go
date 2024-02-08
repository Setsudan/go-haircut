package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gohairdresser/auth"
	"gohairdresser/database"
	"gohairdresser/structs"
)

func AuthentificationRoutes(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/client_login", clientLogin)
		r.Post("/saloon_login", saloonLogin)
		r.Post("/client_signup", clientSignup)
		r.Post("/saloon_signup", saloonSignup)
	})
}

func clientSignup(w http.ResponseWriter, r *http.Request) {
	var client structs.CreateClient
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}

	uid, err := auth.CreateClient(client)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error creating client", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Client created successfully", struct {
		UID string `json:"uid"`
	}{UID: uid}, nil)
}

func saloonSignup(w http.ResponseWriter, r *http.Request) {
	var saloon structs.CreateSaloon
	err := json.NewDecoder(r.Body).Decode(&saloon)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}

	uid, err := auth.CreateSaloon(saloon)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error creating saloon", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon created successfully", struct {
		UID string `json:"uid"`
	}{UID: uid}, nil)
}

func clientLogin(w http.ResponseWriter, r *http.Request) {
	var login structs.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}

	_, token, err := auth.LoginClient(login.Email, login.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "Error logging in"

		if err == database.ErrAccountNotFound {
			status = http.StatusNotFound
			message = "Account not found"
		} else if err == database.ErrInvalidPassword {
			status = http.StatusUnauthorized
			message = "Invalid password"
		}

		SendResponse(w, status, "Error", message, nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Login successful", map[string]string{"token": token}, nil)
}

func saloonLogin(w http.ResponseWriter, r *http.Request) {
	var login structs.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}

	_, token, err := auth.LoginAsSaloon(login.Email, login.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "Error logging in"

		if err == database.ErrAccountNotFound {
			status = http.StatusNotFound
			message = "Account not found"
		} else if err == database.ErrInvalidPassword {
			status = http.StatusUnauthorized
			message = "Invalid password"
		}

		SendResponse(w, status, "Error", message, nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Login successful", map[string]string{"token": token}, nil)
}
