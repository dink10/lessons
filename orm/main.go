package main

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	Name     string `default:"postgres"`
	User     string `default:"admin"`
	Password string `default:"admin"`
}

type User struct {
	Id    int
	Name  string
	Age   int
	Email string
}

func main() {

	// config init
	cfg, err := initConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	// end config init

	// db init
	opts, err := pg.ParseURL(
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
		),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	conn := pg.Connect(opts)
	// end db init

	// insertNewRecord(conn)

	if err := selectRecords(conn); err != nil {
		logrus.Error(err)
	}
}

func selectRecords(conn *pg.DB) error {
	var users []User

	if err := conn.Model(&users).Order("id DESC").Limit(1).Select(); err != nil {
		return err
	}

	//if err := conn.Model(&users).Where("age<?", 29).Select(); err != nil {
	//	return err
	//}

	for _, u := range users {
		logrus.Printf("%d : %s : %d : %s", u.Id, u.Name, u.Age, u.Email)
	}

	return nil
}

func insertNewRecord(conn *pg.DB) {
	u := User{
		Name:  "Alis",
		Age:   15,
		Email: "alis@gmail.com",
	}

	// INSERT INT users VALUES("Alis", 15, "alis@gmail.com");
	res, err := conn.Model(&u).Insert()
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(res.RowsAffected())
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
