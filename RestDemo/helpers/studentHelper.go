package helpers

import (
	"RestDemo/models"
	"log"
)

func CreateStudents(records [][]string) []models.Student {
	students := make([]models.Student, 0)
	for i, record := range records {
		student, err := CreateStudent(record)
		if err != nil {
			log.Printf("Error creating student on record %v", i)
			continue
		}
		students = append(students, student)
	}
	return students
}

func CreateStudent(record []string) (models.Student, error) {
	birthday, err := StringToDate(record[3])
	if err != nil {
		log.Println("Error parsing data")
		return models.Student{}, err
	}

	id, firstName, lastName := record[0], record[1], record[2]
	student := models.Student{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthday,
	}

	return student, nil
}
