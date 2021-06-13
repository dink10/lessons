package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	connUrl := "postgres://admin:admin@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connUrl)
	if err != nil {
		logrus.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	path, _ := os.Getwd()

	fmt.Println(path)

	m, err := migrate.NewWithDatabaseInstance(
		"file://Users/dmitrykologrivov/go/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := m.Steps(1); err != nil {
		logrus.Fatal(err)
	}
}
