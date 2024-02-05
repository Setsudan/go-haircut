package database

import (
	"gohairdresser/structs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func CreateSaloon(saloonData structs.CreateSaloon) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec(`
		INSERT INTO hairSaloon (uid, name, address, email, phone, openingtime, closingtime)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, uid, saloonData.Name, saloonData.Address, saloonData.Email, saloonData.Phone, saloonData.OpeningTime, saloonData.ClosingTime)
	if err != nil {
		log.Printf("failed to create HairSaloon: %v", err)
		return "", err
	}

	return uid, nil
}

func CreateClient(client structs.CreateClient) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO clients (uid, email, age, password)
		VALUES (?, ?, ?, ?)
	`, uid, client.Email, client.Age, client.Password)

	if err != nil {
		log.Printf("failed to create Client: %v", err)
		return "", err
	}

	return uid, nil
}

func CreateHairdresser(hairdresserData structs.CreateHairdresser) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO hairdresser (uid, saloonId, firstName, speciality)
		VALUES (?, ?, ?, ?)
	`, uid, hairdresserData.SaloonID, hairdresserData.FirstName, hairdresserData.Speciality)

	if err != nil {
		log.Printf("failed to create Hairdresser: %v", err)
		return "", err
	}

	return uid, nil
}

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

func CreateSchedule(scheduleData structs.Schedule) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO schedule (uid, hairdresserId, startHour, endHour, availability)
		VALUES (?, ?, ?, ?, ?)
	`, uid, scheduleData.HairdresserID, scheduleData.StartHour, scheduleData.EndHour, scheduleData.Availability)

	if err != nil {
		log.Printf("failed to create Schedule: %v", err)
		return "", err
	}

	return uid, nil
}

func CreateReservation(reservationData structs.Reservation) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO reservation (uid, saloonId, clientId, hairdresserId, startHour, endHour, status)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, uid, reservationData.SaloonID, reservationData.ClientID, reservationData.HairdresserID, reservationData.StartHour, reservationData.EndHour, reservationData.Status)

	if err != nil {
		log.Printf("failed to create Reservation: %v", err)
		return "", err
	}

	return uid, nil
}
