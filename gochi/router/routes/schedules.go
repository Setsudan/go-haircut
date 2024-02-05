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
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving schedules", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Schedules retrieved successfully", data, nil)
}

func getScheduleByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetScheduleByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Schedule not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Schedule retrieved successfully", data, nil)
}
