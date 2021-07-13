package cor

import "fmt"

type Handler interface {
	SendRequest(message int) string
}

type ConcreteHandlerA struct {
	next Handler
}

func (a *ConcreteHandlerA) SendRequest(message int) string {
	if message == 1 {
		return fmt.Sprintf("It's handler %d", message)
	}

	if a.next == nil {
		return ""
	}

	return a.next.SendRequest(message)
}

type ConcreteHandlerB struct {
	next Handler
}

func (b *ConcreteHandlerB) SendRequest(message int) string {
	if message == 2 {
		return fmt.Sprintf("It's handler %d", message)
	}

	if b.next == nil {
		return ""
	}

	return b.next.SendRequest(message)
}

type ConcreteHandlerC struct {
	next Handler
}

func (c *ConcreteHandlerC) SendRequest(message int) string {
	if message == 3 {
		return fmt.Sprintf("It's handler %d", message)
	}

	if c.next == nil {
		return ""
	}

	return c.next.SendRequest(message)
}
