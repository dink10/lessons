package main

type Animal interface {
	BaseAnimal
	run() int
}

type BaseAnimal interface {
	// volume - ...
	voice()
	eat()
}

func main() {
	cat := Cat{}
	//dog := Dog{}

	var randomAnimal Animal

	randomAnimal = cat
	randomAnimal.voice()

	//saySomething(cat)
	//saySomething(dog)
}

func saySomething(object Animal) {
	object.voice()
}
