package database

import (
	"gohairdresser/structs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

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
		VALUES (?, ?, ?, ?, ?, ?, ?)
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
	`, uid, client.Email, client.Age, client.Password)

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
		INSERT INTO hairdresser (uid, saloonId, firstName, speciality)
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
func CreateAppointments(appointmentsData structs.Appointments) (string, error) {
	uid := uuid.New().String()

	db := SetupDatabase()
	defer db.Close()
	_, err := db.Exec(`
		INSERT INTO appointments (uid, saloonId, clientId, hairdresserId, startHour, status)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, uid, appointmentsData.SaloonID, appointmentsData.ClientID, appointmentsData.HairdresserID, appointmentsData.StartHour, appointmentsData.Status)

	if err != nil {
		log.Printf("failed to create Appointments: %v", err)
		return "", err
	}

	return uid, nil
}
