package flyweight

import (
	"strconv"
	"testing"
)

func TestFlyweight(t *testing.T) {
	testSuite := []struct {
		filename string
		weight   int
		height   int
		opacity  float64
		expected string
	}{
		{
			"cat.jpg", 100, 100, 0.95, "Draw image: cat.jpg, width: 100, height: 100, opacity: 0.95",
		},
		{
			"cat.jpg", 200, 400, 1.00, "Draw image: cat.jpg, width: 200, height: 400, opacity: 1.00",
		},
		{
			"dog.jpg", 56, 75, 0.6, "Draw image: dog.jpg, width: 56, height: 75, opacity: 0.60",
		},
	}

	factory := &FlyweightFactory{}

	for i, tt := range testSuite {
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			flyweight := factory.getFlyweight(tt.filename)

			result := flyweight.Draw(tt.weight, tt.height, tt.opacity)

			if tt.expected != result {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
