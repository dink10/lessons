package builder

type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) MakeHeader(content string) {
	b.product.Content += "<header>" + content + "</header>"
}

func (b *ConcreteBuilder) MakeBody(content string) {
	b.product.Content += "<body>" + content + "</body>"
}

func (b *ConcreteBuilder) MakeFooter(content string) {
	b.product.Content += "<footer>" + content + "</footer>"
}
