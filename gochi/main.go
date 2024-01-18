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
	http.ListenAndServe(":8080", r)
}
