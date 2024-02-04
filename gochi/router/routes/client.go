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
		r.Delete("/{uid}", deleteClientByUID)
	})
}

func getAllClients(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllClients()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving clients", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Clients retrieved successfully", data, nil)
}

func getClientByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetClientByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Client not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Client retrieved successfully", data, nil)
}

func deleteClientByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	err := database.DeleteClient(uid)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error deleting client", nil, err)
		return
	}

	// For successful deletion, you might want to send back an empty object or a confirmation message
	SendResponse(w, http.StatusOK, "Success", "Client deleted successfully", map[string]string{"message": "Client deleted successfully"}, nil)
}
