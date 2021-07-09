package factory_method

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type Creator interface {
	CreateProduct(action action) Product
}
