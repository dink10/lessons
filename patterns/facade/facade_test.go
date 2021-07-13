package facade

import "testing"

func TestFacade(t *testing.T) {
	expected := "House is building:Tree is growing:Child born"

	man := NewMan()

	if expected != man.Todo() {
		t.Errorf("Expected %s, got %s", expected, man.Todo())
	}
}
