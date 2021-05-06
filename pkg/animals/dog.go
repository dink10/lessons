package animals

type Dog struct {
	Animal
	Tail bool
}

func NewDog(name string) Dog {
	return Dog{
		Animal: Animal{
			name: name,
		},
		Tail: false,
	}
}
