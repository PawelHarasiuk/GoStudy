package helpers

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadRecordsFromCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil
	}
	return records
}
