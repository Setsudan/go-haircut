package routes

import (
	"encoding/json"
	"gohairdresser/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ClientsRoutes(r *chi.Mux) {
	r.Route("/clients", func(r chi.Router) {
		r.Get("/all", getAllClients)
		r.Get("/{uid}", getClientByUID)
	})
}

func getAllClients(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllClients(database.SetupDatabase())
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)

	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}

func getClientByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetClientByUID(database.SetupDatabase(), uid)
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)

	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}
