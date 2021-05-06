package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(2)
	go consumer(ch, &wg)
	fmt.Println(runtime.NumGoroutine())
	go producer(ch, &wg)
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		runtime.Gosched()
		fmt.Println("receive", val)
	}

	fmt.Println("consumer finished")
	wg.Done()
}
func producer(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 4; i++ {
		fmt.Println("send", i)
		ch <- i
	}
	close(ch)
	fmt.Println("producer finished")
	wg.Done()
}
