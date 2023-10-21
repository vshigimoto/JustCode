package database

import (
	"database/sql"
	"log"
)

func DB() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(id serial PRIMARY KEY, name text, email text, password text)")
	if err != nil {
		return
	}
}

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
