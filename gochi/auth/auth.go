package auth

import (
	"gohairdresser/database"
	"gohairdresser/structs"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(client structs.Client) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["client_id"] = client.UID
	claims["email"] = client.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func CreateClient(client structs.CreateClient) (string, error) {
	hashedPassword, err := database.HashPassword(client.Password)
	if err != nil {
		return "", err
	}
	client.Password = hashedPassword
	return database.CreateClient(client)
}

func LoginClient(email, password string) (structs.Client, string, error) {
	acc, err := database.LoginClient(email, password)
	if err != nil {
		// You can use custom error types or error messages to differentiate the errors
		if err == database.ErrAccountNotFound {
			return structs.Client{}, "", err
		} else if err == database.ErrInvalidPassword {
			return structs.Client{}, "", err
		}
		// Handle other possible errors
		return structs.Client{}, "", err
	}

	// Generate JWT token for the authenticated client
	token, err := GenerateJWT(acc)
	if err != nil {
		// Handle JWT generation error
		return structs.Client{}, "", err
	}

	return acc, token, nil
}

func DeleteClient(uid string) error {
	return database.DeleteClient(uid)
}

func CreateSaloon(saloon structs.CreateSaloon) (string, error) {
	return database.CreateSaloon(saloon)
}
