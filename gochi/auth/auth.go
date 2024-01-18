package auth

/* In this file, we will create a middleware for the authentification of the user.
 * We will use the middleware provided by chi to create a middleware that will
 * check if the user is authenticated or not.
 * We will also have a sign in and sign up route. As well as a sign out route.
 * We will also have the functions related to them.
 */

/* import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
) */

// Auth is the struct that will hold the information about the user.

// There's three possible roles
// Saloon, Clients, Admin, Hairdresser

// Structs
type Auth struct {
	UID      string
	Email    string
	Password string
	Role     string
}
