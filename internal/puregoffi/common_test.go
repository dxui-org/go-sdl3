package puregoffi

import "testing"

func TestBoolToUintptr(t *testing.T) {
	if BoolToUintptr(false) != 0 || BoolToUintptr(true) != 1 {
		t.Fatal("unexpected boolean conversion")
	}
}
