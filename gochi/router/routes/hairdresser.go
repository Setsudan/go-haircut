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

func HairdresserRoutes(r *chi.Mux) {
	r.Route("/hairdressers", func(r chi.Router) {
		r.Get("/all", getAllHairdressers)
		r.Get("/{uid}", getHairdresserByUID)
		r.Post("/create", createHairdresser)
		r.Post("/isAvailable", isHairdresserAvailable)
		r.Delete("/delete/{uid}", deleteHairdresser)
		//r.Put("/update/{uid}", updateHairdresser)
	})
}

func getAllHairdressers(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllHairdressers()
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error retrieving hairdressers", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdressers retrieved successfully", data, nil)
}

func getHairdresserByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	data, err := database.GetHairdresserByUID(uid)
	if err != nil {
		SendResponse(w, http.StatusNotFound, "Error", "Hairdresser not found", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdresser retrieved successfully", data, nil)
}

func createHairdresser(w http.ResponseWriter, r *http.Request) {
	var hairdresser structs.CreateHairdresser
	err := json.NewDecoder(r.Body).Decode(&hairdresser)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}

	uid, err := database.CreateHairdresser(hairdresser)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error creating hairdresser", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdresser created successfully", struct {
		UID string `json:"uid"`
	}{UID: uid}, nil)
}

/*
	 func updateHairdresser(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		var hairdresser database.UpdateHairdresser
		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
			return
		}

		err = database.UpdateHairdresser(uid, hairdresser)
		if err != nil {
			SendResponse(w, http.StatusInternalServerError, "Error", "Error updating hairdresser", nil, err)
			return
		}

		SendResponse(w, http.StatusOK, "Success", "Hairdresser updated successfully", nil, nil)
	}
*/

func deleteHairdresser(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	err := database.DeleteHairdresser(uid)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error deleting hairdresser", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdresser deleted successfully", nil, nil)
}

func isHairdresserAvailable(w http.ResponseWriter, r *http.Request) {
	type HairdresserAvailability struct {
		HairdresserID   string `json:"hairdresserID"`
		StartHour       string `json:"startHour"`
		AppointmentDate string `json:"appointmentDate"`
	}

	var data HairdresserAvailability
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Error", "Invalid request payload", nil, err)
		return
	}
	// Parse appointment date without the time component
	layout := "2006-01-02" // Date layout without time component
	appointmentDate, err := time.Parse(layout, data.AppointmentDate)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error parsing appointment date", nil, err)
		return
	}

	// Log parsed appointment date for debugging
	fmt.Println("Parsed appointment date:", appointmentDate)

	// Check the year parsed from the date
	fmt.Println("Parsed year:", appointmentDate.Year())

	available, err := database.IsHairdresserAvailable(data.HairdresserID, appointmentDate, appointmentDate)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, "Error", "Error checking hairdresser availability", nil, err)
		return
	}

	SendResponse(w, http.StatusOK, "Success", "Hairdresser availability checked successfully", struct {
		Available bool `json:"available"`
	}{Available: available}, nil)
}
