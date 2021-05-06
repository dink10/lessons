package animals

import "fmt"

type Cat struct {
	name string
	Animal
	Pows bool
}

func NewCat(name string) Cat {
	return Cat{
		name: name,
		Animal: Animal{
			name: name,
		},
		Pows: true,
	}
}

func (c *Cat) SaySomething() {
	fmt.Printf("%s said something\n", c.Animal.name)
}

func (c *Cat) goToShop() {
	fmt.Println("Go to shop")
}

func (c *Cat) Info() {
	fmt.Printf("%s has height: %d, weight: %d\n", c.name, c.height, c.weight)
	c.goToShop()
}
