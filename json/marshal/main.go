package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type car struct {
	//id int
	ID       int    `json:"-"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	MaxSpeed int    `json:"maxSpeed,omitempty"`
	Length   int    `json:"length"`
	Width    int    `json:"width"`
}

func (c car) ChangeModel(model string) {
	c.Model = model
}

func main() {
	c := car{
		ID:    1,
		Brand: "BMW",
		Model: "5",
		//MaxSpeed: 270,
		Length: 400,
		Width:  200,
	}

	f, err := os.Create("c.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	enc := json.NewEncoder(f)
	if err := enc.Encode(c); err != nil {
		log.Fatal(err)
	}

	f.Seek(0, io.SeekStart)

	var boughtCar car
	if err := json.NewDecoder(f).Decode(&boughtCar); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", boughtCar)

	//c.ChangeModel("Volvo")
	//
	//res, err := json.Marshal(c)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(res)
}
