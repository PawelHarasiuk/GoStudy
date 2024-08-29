package repositories

import (
	"TodoApp/types"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type PostgresRepository struct {
	DriverName string
	ConnString string
}

func (rep *PostgresRepository) GetTasks() ([]types.Todo, error) {
	todos := make([]types.Todo, 0)
	conn, err := sql.Open(rep.DriverName, rep.ConnString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	query := "SELECT * FROM todos"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo types.Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.DueDate, &todo.IsCompleted)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// TODO assert id is unique and user cant add it, it happens auto
func (rep *PostgresRepository) CreateTask(newTask *types.Todo) error {
	query := "INSERT INTO todos (id, title, description, dueDate, isCompleted) (VALUES($1, $2, $3, $4, $5));"
	err := rep.executeQuery(query, newTask.Id, newTask.Title, newTask.Description, time.Time{}, newTask.IsCompleted)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepository) UpdateTask(newTask *types.Todo) error {
	query := "UPDATE todos SET title=$1, description=$2, dueTime=$3, isCompleted=$4 WHERE id=$5;"
	err := rep.executeQuery(query, newTask.Title, newTask.Description, time.Time{}, newTask.IsCompleted, newTask.Id)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepository) DeleteTask(id int) error {
	query := "DELETE FROM todos WHERE id=$1;"
	err := rep.executeQuery(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepository) CompleteTask(id int) error {
	query := "UPDATE todos SET isCompleted=true WHERE id=$1;"
	err := rep.executeQuery(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepository) UnCompleteTask(id int) error {
	query := "UPDATE todos SET isCompleted=false WHERE id=$1;"
	err := rep.executeQuery(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepository) executeQuery(query string, params ...any) error {
	conn, err := sql.Open(rep.DriverName, rep.ConnString)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}
