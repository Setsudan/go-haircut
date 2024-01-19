package structs

import (
	"time"
)

type Client struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}

type CreateHairSaloon struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
}

type HairSaloon struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
}

type Hairdresser struct {
	UID        string `json:"uid"`
	SaloonID   string `json:"saloonId"`
	FirstName  string `json:"firstName"`
	Speciality string `json:"speciality"`
}

type Schedule struct {
	UID           string    `json:"uid"`
	HairdresserID string    `json:"hairdresserId"`
	StartHour     time.Time `json:"startHour"`
	EndHour       time.Time `json:"endHour"`
	Availability  bool      `json:"availability"`
}

type Reservation struct {
	UID           string    `json:"uid"`
	SaloonID      string    `json:"saloonId"`
	ClientID      string    `json:"clientId"`
	HairdresserID string    `json:"hairdresserId"`
	StartHour     time.Time `json:"startHour"`
	EndHour       time.Time `json:"endHour"`
	Status        string    `json:"status"`
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
