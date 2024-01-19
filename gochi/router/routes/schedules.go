package routes

import (
	"encoding/json"
	"gohairdresser/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SchedulesRoutes(r *chi.Mux) {
	r.Route("/schedules", func(r chi.Router) {
		r.Get("/all", getAllSchedules)
		r.Get("/{uid}", getScheduleByUID)
	})
}

func getAllSchedules(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllSchedules(database.SetupDatabase())
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}

func getScheduleByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetScheduleByUID(database.SetupDatabase(), uid)
	if err != nil {
		log.Println(err)
	}

	res, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(res)
}
