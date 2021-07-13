package decorator

import (
	"log"
	"sync"
	"testing"
)

func TestDecorator(t *testing.T) {
	expected := "<span>It's component</span>"

	decorator := &ConcreteDecorator{component: &ConcreteComponent{}}

	if expected != decorator.Operation() {
		t.Errorf("Expected %s, got %s", expected, decorator.Operation())
	}

	ms := MyStruct{}
	ms.Lock()
	ms.Unlock()
}

type MyStruct struct {
	mux sync.Mutex
}

func (ms *MyStruct) Lock() {
	log.Println("Lock started")
	ms.mux.Lock()
}

func (ms *MyStruct) Unlock() {
	log.Println("Lock finished")
	ms.mux.Unlock()
}
