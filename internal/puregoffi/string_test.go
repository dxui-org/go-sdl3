//go:build windows || linux || darwin

package puregoffi

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestStringConversions(t *testing.T) {
	for _, value := range []string{"", "SDL 世界"} {
		ptr := BytePtrFromString(value)
		if got := BytePtrToString(ptr); got != value {
			t.Fatalf("round trip: got %q, want %q", got, value)
		}
		bytes := unsafe.Slice(ptr, len(value)+1)
		if bytes[len(value)] != 0 {
			t.Fatalf("%q is not NUL terminated", value)
		}
		runtime.KeepAlive(ptr)
	}
	if got := BytePtrToString(nil); got != "" {
		t.Fatalf("nil pointer: got %q", got)
	}
}

func TestBytePtrFromStringRejectsEmbeddedNUL(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for embedded NUL")
		}
	}()
	BytePtrFromString("SDL\x00suffix")
}
