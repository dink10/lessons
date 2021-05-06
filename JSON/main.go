package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type jsonExample struct {
	LastVersion float64  `json:"version"`
	IDs         []int    `json:"ids"`
	Setting     settings `json:"setting"`
	People      []people `json:"people"`
}

type settings struct {
	Auto bool `json:"auto"`
}

type people struct {
	Name      string  `json:"name"`
	Surname   string  `json:"surname"`
	Birthday  string  `json:"birthday"`
	Followers int     `json:"followers"`
	Rating    float64 `json:"rating"`
	Active    bool    `json:"active"`
}

func main() {
	f, err := os.Open("./example.json")
	logError(err)

	data, err := io.ReadAll(f)
	logError(err)

	je := jsonExample{}
	logError(json.Unmarshal(data, &je))

	fmt.Println(je)

	var names []string
	for i := range je.People {
		names = append(names, je.People[i].Name)
	}

	spew.Dump(names)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
