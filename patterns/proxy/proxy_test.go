package proxy

import "testing"

func TestProxy(t *testing.T) {
	expected := "<html>Sent in the real subject</html>"

	proxy := &Proxy{}

	if expected != proxy.Send() {
		t.Errorf("Expeced %s, got %s", expected, proxy.Send())
	}
}
