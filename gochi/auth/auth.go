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

func GenerateSaloonJWT(saloon structs.HairSaloon) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["saloon_id"] = saloon.UID
	claims["email"] = saloon.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func GetIdFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["client_id"].(string), nil
	}
	return "", err
}

// //////////////////////
// Client operations //
// ////////////////////
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

////////////////////////////
// Hair Saloon operations //
//////////////////////////

func CreateSaloon(saloon structs.CreateSaloon) (string, error) {
	return database.CreateSaloon(saloon)
}

func LoginAsSaloon(email, password string) (structs.HairSaloon, string, error) {
	acc, err := database.LoginSaloon(email, password)
	if err != nil {
		if err == database.ErrAccountNotFound {
			return structs.HairSaloon{}, "", err
		} else if err == database.ErrInvalidPassword {
			return structs.HairSaloon{}, "", err
		}
		return structs.HairSaloon{}, "", err
	}

	token, err := GenerateSaloonJWT(acc)
	if err != nil {
		return structs.HairSaloon{}, "", err
	}

	return acc, token, nil
}

/* func DeleteSaloon(uid string) error {
	return database.DeleteSaloon(uid)
} */
