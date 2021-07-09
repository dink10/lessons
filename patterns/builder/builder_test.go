package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	expectedResult := "<header>Header</header><body>Body</body><footer>Footer</footer>"

	product := Product{}

	director := Director{&ConcreteBuilder{product: &product}}
	director.Construct()

	if product.Show() != expectedResult {
		t.Errorf("Expected result %s, got %s", expectedResult, product.Show())
	}
}
