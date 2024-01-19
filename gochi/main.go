package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"gohairdresser/notification"
)

// Res struct
type Res struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// CORS
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))

	serverStatusRes := Res{
		Code:    200,
		Status:  "OK",
		Message: "Server is running",
		Data:    nil,
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serverStatusRes)
	})

	err := notification.SendEmail(notification.EmailParams{
		ToEmail:   "lny.eth@gmail.com",
		Subject:   "RDV accepté",
		HTMLFile:  "./notification/mail_content.gohtml",
		Name:      "Goloum",
		Date:      "22 janvier 2024",
		StartHour: "10:00",
		EndHour:   "11:00",
	})
	if err != nil {
		log.Fatalf("Fail to send email %s", err)
	}

	http.ListenAndServe(":8080", r)
}
