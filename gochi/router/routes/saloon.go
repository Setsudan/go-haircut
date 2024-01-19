package routes

import (
	"encoding/json"
	"gohairdresser/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaloonRoutes(r *chi.Mux) {
	r.Route("/saloons", func(r chi.Router) {
		r.Get("/all", getAllHairSaloons)
		r.Get("/{uid}", getHairSaloonByUID)
	})
}

func getAllHairSaloons(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllHairSaloons(database.SetupDatabase())
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)

	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}

func getHairSaloonByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairSaloonByUID(database.SetupDatabase(), uid)
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)

	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}
