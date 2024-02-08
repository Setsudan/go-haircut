package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteClient(uid string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec(`
		DELETE FROM clients WHERE uid = ?
	`, uid)
	if err != nil {
		log.Printf("failed to delete client: %v", err)
		return err
	}

	return nil
}

func DeleteSaloon(uid string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec(`
		DELETE FROM hairSaloon WHERE uid = ?
	`, uid)
	if err != nil {
		log.Printf("failed to delete saloon: %v", err)
		return err
	}

	return nil
}

func DeleteHairdresser(uid string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec(`
		DELETE FROM hairdressers WHERE uid = ?
	`, uid)
	if err != nil {
		log.Printf("failed to delete hairdresser: %v", err)
		return err
	}

	return nil
}

func DeleteAllAppointments() error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec(`
		DELETE FROM appointments
	`)
	if err != nil {
		log.Printf("failed to delete appointments: %v", err)
		return err
	}

	return nil
}
