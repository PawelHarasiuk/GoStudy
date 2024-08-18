package helpers

import "database/sql"
import _ "github.com/lib/pq"

func ExecuteQuery(query string, connString string, params ...any) error {
	conn, err := sql.Open("postgres", connString)
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
