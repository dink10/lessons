package main

import (
	"fmt"

	"github.com/patrickmn/go-cache"
)

type myInterface interface{}

// itab - (string) ..., data - "string"

type T1 struct{}

func (T1) t1() {
}

func (T1) t11() {
}

type T2 struct{}

func (T2) t2() {
}

type T3 struct{}

func (T3) t3() {
}

func main() {
	//display(T1{})
	//display(T2{})
	//display(T3{})

	c := cache.New(cache.DefaultExpiration, cache.NoExpiration)

	c.Add("key1", 10, cache.DefaultExpiration)
	c.Add("key2", "Hello", cache.DefaultExpiration)

	k1, _ := c.Get("key1")
	fmt.Println(k1.(string))
}

var result = 0

// int - int32/int64

func sum(item interface{}) {
	switch item.(type) {
	case int, int64, int32:
		result += item.(int)
	case float64:
		result += int(item.(float64))
	default:
		fmt.Println("Unknown type")
	}
}

func display(item myInterface) {
	switch x := item.(type) {
	case T1:
		x.t1()
		x.t11()
	case T2:
		x.t2()
	case T3:
		x.t3()
	default:
		fmt.Println("Unknown type")
	}
}
