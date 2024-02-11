package database

import (
	"errors"
	"fmt"
	"gohairdresser/notification"
	"gohairdresser/structs"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var ErrHairdresserNotAvailable = errors.New("hairdresser not available")

/*
 * CreateSaloon creates a new saloon in the database
 * @param saloonData structs.CreateSaloon
 * @return string, error
 */
func CreateSaloon(saloonData structs.CreateSaloon) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec(`
		INSERT INTO hairSaloon (uid, name, address, email, phone, openingtime, closingtime, password)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, uid, saloonData.Name, saloonData.Address, saloonData.Email, saloonData.Phone, saloonData.OpeningTime, saloonData.ClosingTime, saloonData.Password)
	if err != nil {
		log.Printf("failed to create HairSaloon: %v", err)
		return "", err
	}

	return uid, nil
}

/*
 * CreateClient creates a new client in the database
 * @param client structs.CreateClient
 * @return string, error
 */
func CreateClient(client structs.CreateClient) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO clients (uid, email, age, password)
		VALUES (?, ?, ?, ?)
	`, uid, client.Email, 20, client.Password)

	if err != nil {
		log.Printf("failed to create Client: %v", err)
		return "", err
	}

	return uid, nil
}

/*
 * CreateHairdresser creates a new hairdresser in the database
 * @param hairdresserData structs.CreateHairdresser
 * @return string, error
 */
func CreateHairdresser(hairdresserData structs.CreateHairdresser) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO hairdressers (uid, saloonId, firstName, speciality)
		VALUES (?, ?, ?, ?)
	`, uid, hairdresserData.SaloonID, hairdresserData.FirstName, hairdresserData.Speciality)

	if err != nil {
		log.Printf("failed to create Hairdresser: %v", err)
		return "", err
	}

	return uid, nil
}

/*
 * CreateAdmin creates a new admin in the database
 * @param adminData structs.CreateAdmin
 * @return string, error
 */
func CreateAdmin(adminData structs.CreateAdmin) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO admin (uid, name, email, password)
		VALUES (?, ?, ?, ?)
	`, uid, adminData.Name, adminData.Email, adminData.Password)

	if err != nil {
		log.Printf("failed to create Admin: %v", err)
		return "", err
	}

	return uid, nil
}

/*
 * CreateAppointments creates a new appointments in the database
 * @param appointmentsData structs.Appointments
 * @return string, error
 */
func CreateAppointment(appointmentsData structs.CreateAppointment) (string, error) {
	fmt.Print("Database CreateAppointment")
	uid := uuid.New().String()
	status := "Booked"

	startHourTime, err := time.Parse("15:04", appointmentsData.StartHour)
	if err != nil {
		log.Printf("failed to parse start hour: %v", err)
		return "", err
	}

	formattedStartHour := startHourTime.Format("15:04:05") // Assuming the format expected by SQL is "15:04:05"

	db := SetupDatabase()
	defer db.Close()
	_, err = db.Exec(`
    INSERT INTO appointments (uid, saloonId, clientId, hairdresserId, startHour, status, appointmentDate)
    VALUES (?, ?, ?, ?, ?, ?, ?)
`, uid, appointmentsData.SaloonID, appointmentsData.ClientID, appointmentsData.HairdresserID, formattedStartHour, status, appointmentsData.AppointmentsDate)

	if err != nil {
		log.Printf("failed to create Appointment: %v", err)
		return "", err
	}

	clientMail, err := GetClientEmail(appointmentsData.ClientID)
	if err != nil {
		log.Printf("failed to get client email: %v", err)
		return "", err
	}
	appointmentStartTime, parseErr := time.Parse("15:04", appointmentsData.StartHour)
	if parseErr != nil {
		log.Printf("failed to parse appointment start time: %v", parseErr)
		return "", parseErr
	}
	appointmentEndHour := appointmentStartTime.Add(time.Hour).Format("15:04")
	appointmentDate := appointmentStartTime.Format("02 janvier 2006")
	saloonName, err := GetSaloonName(appointmentsData.SaloonID)
	if err != nil {
		log.Printf("failed to get saloon name: %v", err)
		return "", err
	}
	saloonAdress, err := GetSaloonAdress(appointmentsData.SaloonID)
	if err != nil {
		log.Printf("failed to get saloon address: %v", err)
		return "", err
	}
	// Send notification to the client
	notification.SendEmail(notification.EmailParams{
		ToEmail:       clientMail,
		Subject:       "RDV accepté",
		Date:          appointmentDate,
		StartHour:     appointmentsData.StartHour,
		EndHour:       appointmentEndHour,
		SaloonName:    saloonName,
		Description:   `Rendez vous de coiffure, le {{.Date}} à {{.StartHour}} pour une durée d'une heure.`,
		SaloonAddress: saloonAdress,
	})

	return uid, nil
}
