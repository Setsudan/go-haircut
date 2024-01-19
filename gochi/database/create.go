package database

import (
	"gohairdresser/structs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func CreateSaloon(saloonData structs.CreateHairSaloon) (string, error) {
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
