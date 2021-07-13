package decorator

type Component interface {
	Operation() string
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() string {
	return "It's component"
}

type ConcreteDecorator struct {
	component Component
}

func (d *ConcreteDecorator) Operation() string {
	return "<span>" + d.component.Operation() + "</span>"
}
