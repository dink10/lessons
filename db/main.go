package main

import (
	"database/sql"
	"fmt"

	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"

	//mysql
	_ "github.com/go-sql-driver/mysql"

	// Postgres derive
	_ "github.com/lib/pq"
)

const driver = "postgres"

type config struct {
	DB struct {
		Host     string `default:"localhost"`
		Port     int    `default:"5432"`
		Name     string `default:"postgres"`
		User     string `default:"admin"`
		Password string `default:"admin"`
	}
}

type student struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	cfg, err := initConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	url := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", driver, cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	conn, err := sql.Open(driver, url)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	res, err := conn.Query("SELECT id, first_name, last_name FROM students;")
	if err != nil {
		logrus.Fatal(err)
	}

	var students []student
	for res.Next() {
		var stdt student

		if err := res.Scan(&stdt.ID, &stdt.FirstName, &stdt.LastName); err != nil {
			logrus.Fatal(err)
		}

		students = append(students, stdt)
	}

	for _, v := range students {
		fmt.Printf("Student #%d %s:%s\n\n", v.ID, v.FirstName, v.LastName)
	}

	fmt.Println("Gracefully stopped")
}

func initConfig() (*config, error) {
	// Create a new DefaultLoader without or with an initial config file
	m := multiconfig.New()

	// Get an empty struct for your configuration
	var cfg config

	// Populated the serverConf struct
	err := m.Load(&cfg) // Check for error
	if err != nil {
		return nil, err
	}
	m.MustLoad(&cfg) // Panic's if there is any error

	return &cfg, err
}
