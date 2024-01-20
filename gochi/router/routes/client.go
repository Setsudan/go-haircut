package routes

import (
	"gohairdresser/database"
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
	data, err := database.GetAllClients()
	if err != nil {
		SendErrorResponse(w, "Error retrieving clients", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getClientByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetClientByUID(uid)
	if err != nil {
		SendErrorResponse(w, "Client not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
