package structs

import (
	"time"
)

type CreateClient struct {
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type Client struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSaloon struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
	Password    string `json:"password"`
}

type HairSaloon struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
	Password    string `json:"password"`
}

type CreateHairdresser struct {
	SaloonID   string `json:"saloonId"`
	FirstName  string `json:"firstName"`
	Speciality string `json:"speciality"`
}

type Hairdresser struct {
	UID        string `json:"uid"`
	SaloonID   string `json:"saloonId"`
	FirstName  string `json:"firstName"`
	Speciality string `json:"speciality"`
}

type Appointments struct {
	UID           string    `json:"uid"`
	SaloonID      string    `json:"saloonId"`
	ClientID      string    `json:"clientId"`
	HairdresserID string    `json:"hairdresserId"`
	StartHour     time.Time `json:"startHour"`
	Status        string    `json:"status"`
}

type CreateAppointment struct {
	SaloonID      string    `json:"saloonId"`
	ClientID      string    `json:"clientId"`
	HairdresserID string    `json:"hairdresserId"`
	StartHour     time.Time `json:"startHour"`
}

type CreateAdmin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Admin struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type APIResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
