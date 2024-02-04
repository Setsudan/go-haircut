package routes

import (
	"gohairdresser/database"
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
	data, err := database.GetAllSchedules()
	if err != nil && data == nil {
		SendErrorResponse(w, "Error retrieving schedules", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getScheduleByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetScheduleByUID(uid)
	if err != nil {
		SendErrorResponse(w, "Schedule not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
