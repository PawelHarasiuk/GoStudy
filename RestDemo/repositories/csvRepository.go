package repositories

import (
	"RestDemo/models"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type CsvRepository struct {
	Path string
}

func (repository CsvRepository) Load() map[string]models.Student {
	students := make(map[string]models.Student)
	open, err := os.Open(repository.Path)
	if err != nil {
		return nil
	}
	reader := csv.NewReader(open)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		id, firstName, lastName := record[0], record[1], record[2]
		birthday, err := time.Parse("02.01.2006 15:04:05", record[3])
		if err != nil {
			fmt.Printf("error %v", err)
		}
		student := models.Student{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
			Birthdate: birthday,
		}
		students[id] = student
	}
	return students
}

func (repository CsvRepository) Save(students map[string]models.Student) {
	open, err := os.OpenFile(repository.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer open.Close()
	writer := csv.NewWriter(open)
	for id, student := range students {
		record := []string{
			id,
			student.FirstName,
			student.LastName,
			student.Birthdate.Format("02.01.2006 15:04:05"),
		}
		err := writer.Write(record)
		writer.Flush()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
