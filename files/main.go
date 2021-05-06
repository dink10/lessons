package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var list = map[string]int{}

func main() {
	ch := make(chan string, 10)

	m := sync.Map{}
	m.Load("fff")

	var wg sync.WaitGroup
	wg.Add(2)
	var mux sync.Mutex
	go reader(ch, &wg, &mux)
	go writer(ch, &wg, &mux)

	wg.Wait()
}

func reader(ch chan<- string, wg *sync.WaitGroup, mux *sync.Mutex) {
	f, err := os.Open("./files/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range content {
		mux.Lock()
		list[string(v)]++
		mux.Unlock()
		fmt.Println("Send: ", string(v))
		ch <- string(v)
	}

	close(ch)

	wg.Done()
}

func writer(ch <-chan string, wg *sync.WaitGroup, mux *sync.Mutex) {
	f, err := os.Create("files/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		mux.Lock()
		list[v]++
		mux.Unlock()

		fmt.Println("Received: ", v)

		_, err := f.WriteString(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Done()
}
