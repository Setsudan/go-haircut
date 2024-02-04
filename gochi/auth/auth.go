package auth

import (
	"gohairdresser/database"
	"gohairdresser/structs"
)

func CreateClient(client structs.CreateClient) (string, error) {
	hashedPassword, err := database.HashPassword(client.Password)
	if err != nil {
		return "", err
	}
	client.Password = hashedPassword
	return database.CreateClient(client)
}

func DeleteClient(uid string) error {
	return database.DeleteClient(uid)
}

func CreateSaloon(saloon structs.CreateSaloon) (string, error) {
	return database.CreateSaloon(saloon)
}
