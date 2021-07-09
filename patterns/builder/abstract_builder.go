package builder

type Builder interface {
	MakeHeader(content string)
	MakeBody(content string)
	MakeFooter(content string)
}
