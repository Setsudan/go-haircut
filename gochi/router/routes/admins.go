package routes

import (
	"gohairdresser/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AdminsRoutes(r *chi.Mux) {
	r.Route("/admins", func(r chi.Router) {
		r.Get("/all", getAllAdmins)
		r.Get("/{uid}", getAdminByUID)
	})
}

func getAllAdmins(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllAdmins()
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "No admins found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Admins retrieved successfully", data, nil)
}

func getAdminByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetAdminByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Admin not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Admin retrieved successfully", data, nil)
}
