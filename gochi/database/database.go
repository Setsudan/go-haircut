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

	Tables(db)
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

func Tables(db *sql.DB) {
	fakeUser := structs.Client{
		UID:      "uid1",
		Email:    "emaildetest@gmail.com",
		Age:      23,
		Password: "azert",
	}
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS clients (uid VARCHAR(255) PRIMARY KEY, email VARCHAR(255) NOT NULL, age INT NOT NULL, password VARCHAR(255) NOT NULL)")
	if err != nil {
		log.Fatalf("failed to create table %v", err)
	}
	/*
		inject test user
	*/
	_, err = db.Exec("INSERT INTO clients (uid, email, age, password) VALUES (?, ?, ?, ?)", fakeUser.UID, fakeUser.Email, fakeUser.Age, fakeUser.Password)
	if err != nil {
		log.Fatalf("failed to create user %v", err)
	}

	fakeHairdresser := structs.Hairdresser{
		UID:        "uid2",
		SaloonID:   "1",
		FirstName:  "John",
		Speciality: "Colorist",
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS hairdressers (
            uid VARCHAR(255) PRIMARY KEY,
            saloonID INT NOT NULL,
            firstName VARCHAR(255) NOT NULL,
            speciality VARCHAR(255) NOT NULL
        )
    `)
	if err != nil {
		log.Fatalf("failed to create hairdressers table %v", err)
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

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS admin (
            uid VARCHAR(255) PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL
        )
    `)
	if err != nil {
		log.Fatalf("failed to create admin table %v", err)
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

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS hairsaloon (
            uid VARCHAR(255) PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            phone VARCHAR(255) NOT NULL,
			openingtime TIMESTAMP,
			closingtime TIMESTAMP,
        )
    `)
	if err != nil {
		log.Fatalf("failed to create hairsaloon table %v", err)
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

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS reservation (
            uid VARCHAR(255) PRIMARY KEY,
            saloonid VARCHAR(255) NOT NULL,
            clientid VARCHAR(255) NOT NULL,
            hairdresserid FOREIGN KEY (saloonid) REFERENCES hairsaloon,
			starthour TIMESTAMP,
			endhour TIMESTAMP,
			status VARCHAR(355) NOT FULL
			
        )
    `)
	if err != nil {
		log.Fatalf("failed to create reservation table %v", err)
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

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS schedules (
            uid VARCHAR(255) PRIMARY KEY,
            hairdresserid FOREIGN KEY (saloonid) REFERENCES hairsaloon,
			starthour TIMESTAMP,
			endhour TIMESTAMP,
			avaibility BOOLEAN
			
        )
    `)
	if err != nil {
		log.Fatalf("failed to create schedules table %v", err)
	}

	_, err = db.Exec("INSERT INTO schedules (uid, hairdresserid, starthour, endhour, avaibility) VALUES (?, ?, ?, ?, ?)", fakeSchedules.UID, fakeSchedules.HairdresserID, fakeSchedules.StartHour, fakeSchedules.EndHour, fakeSchedules.Availability)
	if err != nil {
		log.Fatalf("failed to create schedules %v", err)
	}

}
