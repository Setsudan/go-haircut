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
