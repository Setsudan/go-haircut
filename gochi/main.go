package main

import (
	"fmt"
	"net/http"

	"gohairdresser/database"
	"gohairdresser/router"
)

func main() {
	r := router.SetupRouter()

	db := database.SetupDatabase()
	defer db.Close()

	database.ShowTables(db)

	fmt.Println("Server is running on port 8080")

	/*err := notification.SendEmail(notification.EmailParams{
		ToEmail:   "lny.eth@gmail.com",
		Subject:   "RDV accepté",
		HTMLFile:  "./notification/mail_content.gohtml",
		Name:      "Goloum",
		Date:      "22 janvier 2024",
		StartHour: "10:00",
		EndHour:   "11:00",
	})
	if err != nil {
		log.Fatalf("Fail to send email %s", err)
	}*/

	// Log all routes
	http.ListenAndServe(":8080", r)
}
