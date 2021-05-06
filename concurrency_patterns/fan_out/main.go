package main

import (
	"fmt"
	"time"
)

func main() {
	fanOut()
}

func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	const cap = 5
	sem := make(chan struct{}, cap)

	for i := 1; i <= emps; i++ {
		go func(emp int) {
			sem <- struct{}{}
			{
				time.Sleep(1 * time.Second)
				ch <- "finished task by employee: " + fmt.Sprint(emp)
			}
			<-sem
		}(i)
		time.Sleep(500 * time.Millisecond)
	}

	for emps > 0 {
		res := <-ch
		fmt.Println("manager received result: ", res)
		emps--
	}

	time.Sleep(20 * time.Second)
}
