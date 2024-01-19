package database

import (
	"database/sql"
	"gohairdresser/structs"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SetupDatabase() *sql.DB {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	createTables(db)
	return db
}

func ShowTables(db *sql.DB) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		log.Println(tableName)
	}
}

func checkError(err error, context string) {
	if err != nil {
		log.Fatalf("failed to %s: %v", context, err)
	}
}

func createTables(db *sql.DB) {
	// Client table
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS clients (
				uid VARCHAR(255) PRIMARY KEY,
				email VARCHAR(255) NOT NULL,
				age INT NOT NULL,
				password VARCHAR(255) NOT NULL
			)
		`)
	checkError(err, "clients")

	// Hairdresser table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS hairdressers (
				uid VARCHAR(255) PRIMARY KEY,
				salonID VARCHAR(255) NOT NULL,
				firstName VARCHAR(255) NOT NULL,
				speciality VARCHAR(255) NOT NULL
			)
		`)
	checkError(err, "hairdressers")

	// Admin table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS admin (
				uid VARCHAR(255) PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL
			)
		`)
	checkError(err, "admin")

	// Hair salon table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS hairSalon (
				uid VARCHAR(255) PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				phone VARCHAR(255) NOT NULL,
				openingTime TIMESTAMP,
				closingTime TIMESTAMP
			)
		`)
	checkError(err, "hairSalon")

	// Reservation table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS reservation (
				uid VARCHAR(255) PRIMARY KEY,
				salonID VARCHAR(255) NOT NULL,
				clientID VARCHAR(255) NOT NULL,
				hairdresserID VARCHAR(255),
				startHour TIMESTAMP,
				endHour TIMESTAMP,
				status VARCHAR(255),
				FOREIGN KEY (hairdresserID) REFERENCES hairdressers(uid)
			)
		`)
	checkError(err, "reservation")

	// Schedule table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS schedules (
				uid VARCHAR(255) PRIMARY KEY,
				hairdresserID VARCHAR(255),
				startHour TIMESTAMP,
				endHour TIMESTAMP,
				availability BOOLEAN,
				FOREIGN KEY (hairdresserID) REFERENCES hairdressers(uid)
			)
		`)
	checkError(err, "schedules")

}

func insertTestData(db *sql.DB) {
	fakeUser := structs.Client{
		UID:      "uid1",
		Email:    "emaildetest@gmail.com",
		Age:      23,
		Password: "azert",
	}
	_, err := db.Exec("INSERT INTO clients (uid, email, age, password) VALUES (?, ?, ?, ?)", fakeUser.UID, fakeUser.Email, fakeUser.Age, fakeUser.Password)
	if err != nil {
		log.Fatalf("failed to create user %v", err)
	}

	fakeHairdresser := structs.Hairdresser{
		UID:        "uid2",
		SaloonID:   "1",
		FirstName:  "John",
		Speciality: "Colorist",
	}

	_, err = db.Exec("INSERT INTO hairdressers (uid, saloonID, firstName, speciality) VALUES (?, ?, ?, ?)", fakeHairdresser.UID, fakeHairdresser.SaloonID, fakeHairdresser.FirstName, fakeHairdresser.Speciality)
	if err != nil {
		log.Fatalf("failed to create hairdresser %v", err)
	}

	fakeAdmin := structs.Admin{
		UID:      "uid2",
		Name:     "ethan",
		Email:    "email@gmail.fr",
		Password: "motdepase",
	}

	_, err = db.Exec("INSERT INTO admin (uid, name, email, password) VALUES (?, ?, ?, ?)", fakeAdmin.UID, fakeAdmin.Name, fakeAdmin.Email, fakeAdmin.Password)
	if err != nil {
		log.Fatalf("failed to create admin %v", err)
	}

	fakeHairSaloon := structs.HairSaloon{
		UID:         "uid2",
		Name:        "ethan",
		Email:       "email@gmail.fr",
		Phone:       "0612233456",
		OpeningTime: time.Now(),
		ClosingTime: time.Now(),
	}

	_, err = db.Exec("INSERT INTO hairsaloon (uid, name, email, phone, openingtime, closingtime) VALUES (?, ?, ?, ?, ?, ?)", fakeHairSaloon.UID, fakeHairSaloon.Name, fakeHairSaloon.Email, fakeHairSaloon.Phone, fakeHairSaloon.OpeningTime, fakeHairSaloon.ClosingTime)
	if err != nil {
		log.Fatalf("failed to create hairsaloon %v", err)
	}

	fakeReservation := structs.Reservation{
		UID:           "uid2",
		SaloonID:      "2",
		ClientID:      "3",
		HairdresserID: "4",
		StartHour:     time.Now(),
		EndHour:       time.Now(),
		Status:        "test",
	}

	_, err = db.Exec("INSERT INTO reservation (uid, saloonid, clientid, hairdresserid, starthour, endhour, status) VALUES (?, ?, ?, ?, ?, ?, ?)", fakeReservation.UID, fakeReservation.SaloonID, fakeReservation.ClientID, fakeReservation.HairdresserID, fakeReservation.StartHour, fakeReservation.EndHour, fakeReservation.Status)
	if err != nil {
		log.Fatalf("failed to create reservation %v", err)
	}

	fakeSchedules := structs.Schedule{
		UID:           "uid2",
		HairdresserID: "4",
		StartHour:     time.Now(),
		EndHour:       time.Now(),
		Availability:  true,
	}

	_, err = db.Exec("INSERT INTO schedules (uid, hairdresserid, starthour, endhour, avaibility) VALUES (?, ?, ?, ?, ?)", fakeSchedules.UID, fakeSchedules.HairdresserID, fakeSchedules.StartHour, fakeSchedules.EndHour, fakeSchedules.Availability)
	if err != nil {
		log.Fatalf("failed to create schedules %v", err)
	}

}
