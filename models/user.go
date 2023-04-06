package models

import "database/sql"

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type User struct {
	ID          string
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	LastVersion int    `json:"last_version"`
	Image       string `json:"image"`
}

type Otp struct {
	ID          string
	PhoneNumber string `json:"phone_number"`
	Otp         string `json:"otp"`
}

type NullShortUser struct {
	ID          sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
}

type NullRole struct {
	ID   sql.NullString
	Name sql.NullString
}

type NullExternalID struct {
	ExternalId sql.NullString
}

type NullShorUserPermmison struct {
	ID   sql.NullString
	Name sql.NullString
}
