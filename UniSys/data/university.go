package data

import "time"

type Student struct {
	IndexNumber string    `json:"indexNumber"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Birthdate   time.Time `json:"birthdate"`
	Email       string    `json:"email"`
	MothersName string    `json:"mothersName"`
	FathersName string    `json:"fathersName"`
	Studies     Studies   `json:"studies"`
}

type Studies struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}

type University struct {
	CreatedAt     time.Time       `json:"createdAt"`
	Author        string          `json:"author"`
	Students      []Student       `json:"students"`
	ActiveStudies []ActiveStudies `json:"activeStudies"`
}

type ActiveStudies struct {
	Name             string `json:"name"`
	NumberOfStudents int    `json:"numberOfStudents"`
}
