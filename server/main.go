package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8085", NewRouter()); err != nil {
		log.Fatal(err)
	}
}
