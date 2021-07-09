package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	asserts := []string{"A", "B", "C"}

	factory := NewCreator()

	products := []Product{
		factory.CreateProduct(A),
		factory.CreateProduct(B),
		factory.CreateProduct(C),
	}

	for i, p := range products {
		if action := p.Use(); action != asserts[i] {
			t.Errorf("Expected action: %s, got %s", asserts[i], action)
		}
	}
}
