package bridge

import "testing"

func TestBridge(t *testing.T) {
	expected := "SUZUUUUUKI"

	car := NewCar(&EngineSuzuki{})

	if car.Race() != expected {
		t.Errorf("Expected sound %s, got %s", expected, car.Race())
	}
}
