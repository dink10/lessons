package builder

type Product struct {
	Content string
}

func (p *Product) Show() string {
	return p.Content
}
