package main

import (
	"fmt"
	"time"
)

func main() {
	drop()
}

func drop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for signal := range ch {
			fmt.Println("received signal: ", signal)
		}
	}()

	const work = 20
	for i := 0; i < work; i++ {
		select {
		case ch <- fmt.Sprintf("signal #%d", i):
			fmt.Println("send signal #", i)
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			fmt.Println("skipped signal")
		}
	}

	close(ch)

	time.Sleep(5 * time.Second)
}
