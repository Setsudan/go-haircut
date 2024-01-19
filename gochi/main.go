package main

import (
	"fmt"
	"net/http"

	"gohairdresser/database"
	"gohairdresser/router"
	"encoding/json"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"gohairdresser/notification"
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
	http.ListenAndServe(":8080", r)
}
