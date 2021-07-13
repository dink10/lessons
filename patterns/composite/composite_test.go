package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {
	expected := "/root\n/root/usr\n/root/usr/B\n/root/A\n"
	// /root
	// /root/usr
	// /root/usr/B
	// /root/A

	rootDir := NewDirectory("root")
	usrDir := NewDirectory("usr")
	fileA := NewFile("A")
	fileB := NewFile("B")

	rootDir.Add(usrDir)
	usrDir.Add(fileB)
	rootDir.Add(fileA)

	result := rootDir.Print("")

	if result != expected {
		t.Errorf("Expected result %s, got %s", expected, result)
	}
}
