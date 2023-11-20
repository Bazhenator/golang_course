package db

import (
	"database/sql"
	"fmt"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type MyDatabaseService struct {
	DB Database
}

func New(db Database) MyDatabaseService {
	return MyDatabaseService{DB: db}
}

func (service MyDatabaseService) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get Query: %w", err)
	}
	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	return names, nil
}

func (service MyDatabaseService) SelectUniqueValues(columnName string, tableName string) ([]string, error) {
	query := "SELECT DISTINCT " + columnName + " FROM " + tableName

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get Query: %w", err)
	}
	defer rows.Close()

	var values []string

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}

		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	return values, nil
}
