package models

import "time"

type Student struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthdate time.Time `json:"birthdate"`
}
