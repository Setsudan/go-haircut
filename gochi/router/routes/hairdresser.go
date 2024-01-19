package routes

import (
	"encoding/json"
	"gohairdresser/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HairdresserRoutes(r *chi.Mux) {
	r.Route("/hairdressers", func(r chi.Router) {
		r.Get("/all", getAllHairdressers)
		r.Get("/{uid}", getHairdresserByUID)
		r.Put("/{hairdresserId}/schedule", updateHairdresserSchedule) // as defined previously
	})
}

func getAllHairdressers(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllHairdressers(database.SetupDatabase())
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)

	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}

func getHairdresserByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairdresserByUID(database.SetupDatabase(), uid)
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)

	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}

func updateHairdresserSchedule(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHairdresserSchedule")
}
