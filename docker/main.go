package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	connUrl := "postgres://admin:admin@postgres:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("Hello from app service"))
}
