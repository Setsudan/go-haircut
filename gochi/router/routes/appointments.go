package routes

import (
	"encoding/json"
	"fmt"
	"gohairdresser/database"
	"gohairdresser/structs"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var emptyStruct = struct{}{}

func AppointmentsRoutes(r *chi.Mux) {
	r.Route("/appointments", func(r chi.Router) {
		r.Get("/all", getAllAppointments)
		r.Get("/{uid}", getAppointmentByUID)
		r.Patch("/cancel/{uid}", updateAppointmentsStatusToFalse)
		r.Patch("/saloonName/{uid}", updateSaloonName)
		r.Patch("/saloonAddress/{uid}", updateSaloonAddress)
		r.Patch("/saloonEmail/{uid}", updateSaloonEmail)
		r.Patch("/saloonPhone/{uid}", updateSaloonPhone)
		r.Patch("/saloonOpeningTime/{uid}", updateSaloonOpeningTime)
		r.Patch("/saloonClosingTime/{uid}", updateSaloonClosingTime)
		r.Post("/create", createAppointment)
		r.Delete("/delete/all", deleteAllAppointments)
	})
}

// ===== GET =====
func getAllAppointments(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllAppointments()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving appointments", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointments retrieved successfully", data, nil)
}

func getAppointmentByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetAppointmentsByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Appointment not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointment retrieved successfully", data, nil)
}

// ===== UPDATE =====
func updateAppointmentsStatusToFalse(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	err := database.UpdateAppointmentsStatusToFalse(uid)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating appointment status", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointment status updated successfully", emptyStruct, nil)
}

func updateSaloonName(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	newName := r.FormValue("newName")
	err := database.UpdateSaloonName(uid, newName)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating saloon name", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon name updated successfully", emptyStruct, nil)
}

func updateSaloonAddress(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	newAddress := r.FormValue("newAddress")
	err := database.UpdateSaloonAddress(uid, newAddress)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating saloon address", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon address updated successfully", emptyStruct, nil)

}

func updateSaloonEmail(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	newEmail := r.FormValue("newEmail")
	err := database.UpdateSaloonEmail(uid, newEmail)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating saloon mail", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon email updated successfully", emptyStruct, nil)
}

func updateSaloonPhone(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	newPhone := r.FormValue("newPhone")
	err := database.UpdateSaloonPhone(uid, newPhone)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating saloon phone", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon phone updated successfully", emptyStruct, nil)
}

func updateSaloonOpeningTime(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	newOpeningTime := r.FormValue("newOpeningTime")
	err := database.UpdateSaloonOpeningTime(uid, newOpeningTime)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating saloon opening time", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon opening time updated successfully", emptyStruct, nil)
}

func updateSaloonClosingTime(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	newClosingTime := r.FormValue("newClosingTime")
	err := database.UpdateSaloonClosingTime(uid, newClosingTime)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error updating saloon closing time", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Saloon closing time updated successfully", emptyStruct, nil)
}

// ===== CREATE =====
func createAppointment(w http.ResponseWriter, r *http.Request) {
	type CreateAppointment struct {
		SaloonID         string `json:"saloonId"`
		ClientID         string `json:"clientId"`
		HairdresserID    string `json:"hairdresserId"`
		StartHour        string `json:"startHour"`
		AppointmentsDate string `json:"appointmentDate"`
	}

	var appointment CreateAppointment
	fmt.Println("received request to create appointment")
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}
	// Parse appointment date without the time component
	layout := "2006-01-02"
	appointmentDate, err := time.Parse(layout, appointment.AppointmentsDate)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error parsing appointment date", nil, err)
		return
	}
	fmt.Println("decoded request to create appointment")

	fmt.Print("Trying Database CreateAppointment")
	appointmentData := structs.CreateAppointment{
		SaloonID:         appointment.SaloonID,
		ClientID:         appointment.ClientID,
		HairdresserID:    appointment.HairdresserID,
		StartHour:        appointment.StartHour,
		AppointmentsDate: appointmentDate,
	}
	uid, err := database.CreateAppointment(appointmentData)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error creating appointment", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointment created successfully", struct {
		UID string `json:"uid"`
	}{UID: uid}, nil)
}

// ===== DELETE =====
func deleteAllAppointments(w http.ResponseWriter, r *http.Request) {
	err := database.DeleteAllAppointments()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error deleting appointments", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Appointments deleted successfully", emptyStruct, nil)
}
