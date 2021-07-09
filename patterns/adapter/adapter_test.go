package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adapter := NewAdaptor(&Adaptee{})
	req := adapter.Request()

	if req != "Request" {
		t.Errorf("Epected value %s, got %s", "Request", req)
	}
}
