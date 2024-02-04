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
	if err != nil && data == nil {
		SendErrorResponse(w, "Error retrieving appointments", err, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}

func getAppointmentByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetReservationByUID(uid)
	if err != nil {
		SendErrorResponse(w, "Appointment not found", err, http.StatusNotFound)
		return
	}

	SendJSONResponse(w, data)
}
