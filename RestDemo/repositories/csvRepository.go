package repositories

import (
	"RestDemo/helpers"
	"RestDemo/models"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type CsvRepository struct {
	Path string
}

func (repository CsvRepository) GetStudents() ([]models.Student, error) {
	records := helpers.ReadRecordsFromCSV(repository.Path)
	students := helpers.CreateStudents(records)
	if len(students) == 0 {
		return make([]models.Student, 0), errors.New("students not found")
	}
	return students, nil
}

func (repository CsvRepository) UpdateStudent(newStudent models.Student) error {
	records := helpers.ReadRecordsFromCSV(repository.Path)

	file, err := os.Create(repository.Path)
	if err != nil {
		mess := fmt.Sprintf("Cant find file: %v", err)
		return errors.New(mess)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	seen := false
	for i, record := range records {
		id := record[0]
		if id == newStudent.Id {
			firstName, lastName := newStudent.FirstName, newStudent.LastName
			birthdate := helpers.DateToString(newStudent.Birthdate)
			records[i][1] = firstName
			records[i][2] = lastName
			records[i][3] = birthdate
			seen = true
			break
		}
	}
	if !seen {
		return errors.New("student not found")
	}
	err = writer.WriteAll(records)
	if err != nil {
		mess := fmt.Sprintf("Cant write student to file: %v", err)
		return errors.New(mess)
	}
	return nil
}

func (repository CsvRepository) CreateStudent(newStudent models.Student) error {
	file, err := os.OpenFile(repository.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		mess := fmt.Sprintf("Cant find file: %v", err)
		return errors.New(mess)
	}

	birthdate := helpers.DateToString(newStudent.Birthdate)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		newStudent.Id,
		newStudent.FirstName,
		newStudent.LastName,
		birthdate,
	})

	if err != nil {
		mess := fmt.Sprintf("Cant write student to file: %v", err)
		return errors.New(mess)
	}

	return nil
}

func (repository CsvRepository) DeleteStudent(oldStudent models.Student) error {
	records := helpers.ReadRecordsFromCSV(repository.Path)
	file, err := os.Create(repository.Path)
	if err != nil {
		mess := fmt.Sprintf("Cant find file: %v", err)
		return errors.New(mess)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	seen := false
	for i, record := range records {
		id := record[0]
		if id == oldStudent.Id {
			records = append(records[:i], records[i+1:]...)
			seen = true
			break
		}
	}
	if !seen {
		return errors.New("student not found")
	}
	err = writer.WriteAll(records)
	if err != nil {
		mess := fmt.Sprintf("Cant write student to file: %v", err)
		return errors.New(mess)
	}
	return nil
}
