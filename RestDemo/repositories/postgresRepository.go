package repositories

import (
	"RestDemo/models"
)

type PostgresRepository struct {
	ConnString string
}

func (postgresRepository PostgresRepository) Load() map[string]models.Student {
	//	db, err := sql.Open("postgres", postgresRepository.ConnString)
	//	if err != nil {
	//		fmt.Printf("Problem connecting to db %v", err)
	//	}
	//	return nil
	return nil
}

func (postgresRepository PostgresRepository) Save(map[string]models.Student) {
	//	db, err := sql.Open()
}
