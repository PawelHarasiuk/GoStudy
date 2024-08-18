package repositories

import (
	"RestDemo/helpers"
	"RestDemo/models"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	ConnString string
}

func (repository PostgresRepository) GetStudents() ([]models.Student, error) {
	query := "SELECT * FROM student;"
	var students []models.Student
	conn, err := sql.Open("postgres", repository.ConnString)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	records, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer records.Close()
	for records.Next() {
		var student models.Student
		err := records.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Birthdate)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (repository PostgresRepository) UpdateStudent(newStudent models.Student) error {
	id, firstName, lastName, birthdate := newStudent.Id, newStudent.FirstName, newStudent.LastName, newStudent.Birthdate
	strBirthdate := helpers.DateToString(birthdate)
	query := "UPDATE student SET firstname = $1, lastname = $2, birthdate = $3 WHERE id = $4;"
	err := helpers.ExecuteQuery(query, repository.ConnString, firstName, lastName, strBirthdate, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository PostgresRepository) CreateStudent(newStudent models.Student) error {
	id, firstName, lastName, birthdate := newStudent.Id, newStudent.FirstName, newStudent.LastName, newStudent.Birthdate
	strBirthdate := helpers.DateToString(birthdate)
	query := "INSERT INTO student (id, firstname, lastname, birthdate) (VALUES($1, $2, $3, $4));"
	err := helpers.ExecuteQuery(query, repository.ConnString, id, firstName, lastName, strBirthdate)
	if err != nil {
		return err
	}
	return nil
}

func (repository PostgresRepository) DeleteStudent(oldStudent models.Student) error {
	id := oldStudent.Id
	query := "DELETE FROM student WHERE id = $1"
	err := helpers.ExecuteQuery(query, repository.ConnString, id)
	if err != nil {
		return err
	}
	return nil
}
