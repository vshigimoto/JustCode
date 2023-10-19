package database

import "database/sql"

func Connect() (*sql.DB, error) {
	dataSourceName := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
