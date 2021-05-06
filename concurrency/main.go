package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	pi = 3.14
	mutexStarving2
)

func main() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(pi)
	fmt.Println(mutexStarving2)

	wg := sync.WaitGroup{}
	n := 2
	wg.Add(n)

	var result int32

	for i := 0; i < n; i++ {
		go func(val int) {
			defer wg.Done()
			atomic.AddInt32(&result, 1)
		}(i)
	}

	wg.Wait()

	fmt.Println(result)
}

func sumWithError(x, y int) (int, error) {
	fmt.Println(x, y)

	return x + y, nil
}

func sum(x, y int) int {
	fmt.Println(x, y)

	return x + y
}
