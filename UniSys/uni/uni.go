package uni

import (
	"UniSys/data"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func Export(inputPath, outputPath, format string) {
	records, err := readFile(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	university := createUniversity(records)
	var bytes []byte
	switch format {
	case "json":
		bytes = jsonify(university)
	default:
		fmt.Println("Invalid format")
	}
	saveToFile(outputPath, bytes)
}

func createStudents(records [][]string) []data.Student {
	studentsMap := make(map[string]data.Student)

	for _, record := range records {
		indexNumber, firstName, lastName, email, mothersName, fathersName := record[4], record[0], record[1], record[6], record[7], record[8]
		birthdate, err := time.Parse("2006-01-02", record[5])
		if err != nil {
			fmt.Println("wrong date")
			continue
		}
		studies := data.Studies{
			Name: record[2],
			Mode: record[3],
		}

		index := firstName + "_" + lastName + "_" + indexNumber
		_, present := studentsMap[index]

		if !isRecordValid(record) {
			logError("record invalid " + index)
			continue
		}
		if present {
			logError("student already exists " + index)
			continue
		}

		student := data.Student{
			IndexNumber: indexNumber,
			FirstName:   firstName,
			LastName:    lastName,
			Birthdate:   birthdate,
			Email:       email,
			MothersName: mothersName,
			FathersName: fathersName,
			Studies:     studies,
		}

		studentsMap[index] = student
	}

	students := make([]data.Student, 0)
	for _, student := range studentsMap {
		students = append(students, student)
	}
	return students
}

func createActiveStudies(records [][]string) []data.ActiveStudies {
	countActiveStudies := make(map[string]int)
	activeStudies := make([]data.ActiveStudies, 0)
	for _, record := range records {
		countActiveStudies[record[2]] += 1
	}

	for name, numberOfStudents := range countActiveStudies {
		activeStudies = append(activeStudies, data.ActiveStudies{
			Name:             name,
			NumberOfStudents: numberOfStudents,
		})
	}
	return activeStudies
}

func createUniversity(records [][]string) data.University {
	var wg sync.WaitGroup
	var students []data.Student
	var activeStudies []data.ActiveStudies

	wg.Add(2)
	go func() {
		students = createStudents(records)
		wg.Done()
	}()

	go func() {
		activeStudies = createActiveStudies(records)
		wg.Done()
	}()

	wg.Wait()

	university := data.University{
		CreatedAt:     time.Now(),
		Author:        "Me",
		Students:      students,
		ActiveStudies: activeStudies,
	}
	return university
}

func readFile(inputPath string) ([][]string, error) {
	open, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Wrong path")
		return nil, err
	}

	reader := csv.NewReader(open)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Problem with reading")
		return nil, err
	}

	return records, err
}

func saveToFile(outputPath string, toSave []byte) {
	create, err := os.Create(outputPath + "output.json")
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.WriteString(create, string(toSave))
	if err != nil {
		fmt.Println(err)
	}
}

func isRecordValid(record []string) bool {
	for _, val := range record {
		if val == "" {
			return false
		}
	}
	return true
}

func jsonify(university data.University) []byte {
	marshal, err := json.Marshal(university)
	if err != nil {
		fmt.Println(err)
	}
	return marshal
}

func logError(message string) {
	create, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.WriteString(create, message+"\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}
