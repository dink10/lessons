package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Student struct {
	FIO    string `xml:"fio"`
	Year   string `xml:"year,omitempty"`
	Course int    `xml:"course"`
}

func main() {
	st := &Student{
		FIO:    "Vasily Pupkin",
		Course: 2,
	}

	res, err := xml.Marshal(st)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(res, st)

	fmt.Println(string(res))
}
