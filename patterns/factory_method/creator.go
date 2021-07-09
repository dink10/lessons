package factory_method

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{action: string(action)}
	case B:
		product = &ConcreteProductB{action: string(action)}
	case C:
		product = &ConcreteProductC{action: string(action)}
	}

	return product
}

type ConcreteProductA struct {
	action string
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

type ConcreteProductB struct {
	action string
}

func (p *ConcreteProductB) Use() string {
	return p.action
}

type ConcreteProductC struct {
	action string
}

func (p *ConcreteProductC) Use() string {
	return p.action
}
