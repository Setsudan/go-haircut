package routes

import (
	"encoding/json"
	"gohairdresser/database"
	"log"
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
	data, err := database.GetAllAdmins(database.SetupDatabase())
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}

func getAdminByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetAdminByUID(database.SetupDatabase(), uid)
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}
