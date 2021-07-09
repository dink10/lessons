package adapter

type Target interface {
	Request() string
}

type Adaptee struct {
}

func NewAdaptor(a *Adaptee) Target {
	return &Adapter{Adaptee: a}
}

func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
