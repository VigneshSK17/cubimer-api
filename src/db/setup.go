package db

import (
	"database/sql"
	"fmt"

	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

type DB struct {
	Queries *Queries
}


func Get() (DB, error) {

	q := DB{}

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/New_York",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return q, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return q, err
	}
	
	m, err := migrate.NewWithDatabaseInstance(
		"file:///usr/src/app/db/migrations",
		os.Getenv("DB_NAME"), driver)
	if err != nil {
		log.Print(err.Error())
		return q, err
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		log.Print(err.Error())
		return q, err
	}

	q.Queries = New(db)
	return q, err

}
