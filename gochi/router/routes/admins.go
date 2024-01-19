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
	db := database.SetupDatabase()
	data, err := database.GetAllAdmins(db)
	if err != nil {
		SendErrorResponse(w, "No admins found", err, http.StatusBadRequest)
		return
	}

	SendJSONResponse(w, data)
}

func getAdminByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	db := database.SetupDatabase()
	data, err := database.GetAdminByUID(db, uid)
	if err != nil {
		SendErrorResponse(w, "Admin not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
