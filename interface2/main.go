package main

import (
	"context"
	"os"

	"github.com/dink10/lessons/interface2/db"
)

type Reader interface {
	read() []byte
}

type Writer interface {
	write([]byte)
}

type ReaderWriter interface {
	Reader
	Writer
}

func copy(r Reader, w Writer) {
	b := r.read()
	w.write(b)
}

func main() {
	f, _ := os.Open("")
	defer f.Close()

	driver := MySql{}

	lib := db.DB{DB: driver}
	lib.Save(context.Background(), db.Item{})
}
