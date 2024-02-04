package routes

import (
	"gohairdresser/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AppointmentsRoutes(r *chi.Mux) {
	r.Route("/appointments", func(r chi.Router) {
		r.Get("/all", getAllAppointments)
		r.Get("/{uid}", getAppointmentByUID)
	})
}

func getAllAppointments(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllReservations()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving appointments", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointments retrieved successfully", data, nil)
}

func getAppointmentByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetReservationByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Appointment not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointment retrieved successfully", data, nil)
}
