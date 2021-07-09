package singleton

import "testing"

func TestSingleton(t *testing.T) {
	inst1 := GetInstance()
	inst2 := GetInstance()
	inst3 := GetInstance()
	inst4 := GetInstance()

	if inst1 != inst2 && inst2 != inst3 && inst3 != inst4 && inst1 != inst4 {
		t.Errorf("Objects are not equal!")
	}
}
