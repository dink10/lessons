package main

import (
	"fmt"
	"time"
)

func main() {
	cancellation()
}

func cancellation() {
	ch := make(chan string, 1)

	go func() {
		//time.Sleep(500 * time.Millisecond)
		ch <- "signal"
		fmt.Println("send signal")
	}()

	tc := time.After(100 * time.Millisecond) // now -> now + 100 * time.Millisecond

	select {
	case signal := <-ch:
		fmt.Println("received signal", signal)
	case t := <-tc:
		fmt.Println("cancelled: timeout: ", t)
	}

	// end
}
