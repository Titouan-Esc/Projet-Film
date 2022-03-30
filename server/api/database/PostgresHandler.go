package database

import (
	"database/sql"
	"log"
)

type PostgresHandler struct {
	Conn *sql.DB
}

func (handler *PostgresHandler) Query(statement string) (IRow, error) {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		log.Panicf("Error to make the Query { %v }", err)
		return new(PostgresRow), err
	}

	row := new(PostgresRow)
	row.Rows = rows

	return row, nil
}

type PostgresRow struct {
	Rows *sql.Rows
}

func (handler *PostgresRow) Scan(dest ...interface{}) error {
	if err := handler.Rows.Scan(dest...); err != nil {
		log.Panicf("Error to Scan DB { %v }", err)
		return err
	}

	return nil
}

func (handler *PostgresRow) Next() bool {
	return handler.Rows.Next()
}
