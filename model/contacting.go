package model

import (
// "database/sql"

// "github.com/lib/pq"
)

type Contact struct {
	FirstName   string `json:"" `
	LastName    string `json:""`
	PhoneNumber string `json:""`
	TypeService string `json:""`
	Text        string `json:""`
}

type ContactFull struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	TypeService string `json:"type_service" db:"type_service"`
	Text        string `json:"text" db:"text"`
	Created_At  string `json:"created_at" db:"created_at"`
}
type allContacts struct {
	AllContact []ContactFull
}
