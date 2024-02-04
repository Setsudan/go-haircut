package auth

import (
	"gohairdresser/database"
	"gohairdresser/structs"
)

func CreateClient(client structs.CreateClient) (string, error) {
	client.Password = database.HashPassword(client.Password)
	return database.CreateClient(client)
}

func CreateSaloon(saloon structs.CreateSaloon) (string, error) {
	return database.CreateSaloon(saloon)
}
