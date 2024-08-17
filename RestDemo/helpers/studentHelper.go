package helpers

import (
	"RestDemo/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

var dateFormat = "02.01.2006 15:04:05"

func ReadRecords(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return records
}

func CreateStudents(path string) []models.Student {
	records := ReadRecords(path)
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
		fmt.Println("Error parsing data")
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

func StringToDate(stringDate string) (time.Time, error) {
	birthday, err := time.Parse(dateFormat, stringDate)
	return birthday, err
}

func DateToString(date time.Time) string {
	return time.Time.Format(date, dateFormat)
}
