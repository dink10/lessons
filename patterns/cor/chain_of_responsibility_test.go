package cor

import (
	"fmt"
	"testing"
)

func TestCOR(t *testing.T) {
	testSuite := []struct {
		message  int
		expected string
	}{
		{
			1,
			"It's handler 1",
		},
		{
			2,
			"It's handler 2",
		},
		{
			3,
			"It's handler 3",
		},
		{
			4,
			"",
		},
	}

	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	for i, ts := range testSuite {
		t.Run(fmt.Sprintf("Case: %d", i), func(t *testing.T) {
			if ts.expected != handlers.SendRequest(ts.message) {
				t.Errorf("Expected %s, got %s", ts.expected, handlers.SendRequest(ts.message))
			}
		})
	}
}
