package proxy

type Subject interface {
	Send() string
}

type Proxy struct {
	realSubject Subject
}

func (p *Proxy) Send() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}

	return "<html>" + p.realSubject.Send() + "</html>"
}

type RealSubject struct {
}

func (s *RealSubject) Send() string {
	return "Sent in the real subject"
}
