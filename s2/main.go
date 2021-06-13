package main

import (
	"log"

	"github.com/dink10/lessons/s2/server"
)

func main() {
	s := server.NewServer("localhost", 3000)
	log.Fatal(s.Run())
}
