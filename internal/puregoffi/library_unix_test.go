//go:build linux || darwin

package puregoffi

import (
	"runtime"
	"testing"
)

func TestLibraryLifecycle(t *testing.T) {
	name, symbol := "libc.so.6", "malloc"
	if runtime.GOOS == "darwin" {
		name, symbol = "/usr/lib/libSystem.B.dylib", "malloc"
	}
	handle, err := OpenLibrary(name)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSymbol(handle, symbol); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSymbol(handle, "go_sdl3_symbol_that_does_not_exist"); err == nil {
		t.Fatal("missing symbol unexpectedly resolved")
	}
	if err := CloseLibrary(handle); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenLibrary("go-sdl3-library-that-does-not-exist.so"); err == nil {
		t.Fatal("missing library unexpectedly opened")
	}
}
