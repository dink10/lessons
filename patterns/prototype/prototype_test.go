package prototype

import "testing"

func TestPrototype(t *testing.T) {
	product := NewConcreteProduct("A")

	clone := product.Clone()

	if clone.GetName() != product.GetName() {
		t.Errorf("Expected product name %s, got %s", product.GetName(), clone.Clone())
	}
}
