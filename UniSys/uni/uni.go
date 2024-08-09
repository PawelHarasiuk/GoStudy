package uni

import (
	"UniSys/data"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func Export(inputPath, outputPath, format string) {
	records, err := readFile(inputPath)
	if err != nil {
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
	students := make([]data.Student, 0)
	for _, record := range records {
		if !isRecordValid(record) {
			continue
		}

		student := data.Student{
			IndexNumber: record[4],
			FirstName:   record[0],
			LastName:    record[1],
			Birthdate:   time.Time{},
			Email:       record[6],
			MothersName: record[7],
			FathersName: record[8],
			Studies: data.Studies{
				Name: record[2],
				Mode: record[3],
			},
		}

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
	students := createStudents(records)
	activeStudies := createActiveStudies(records)
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
