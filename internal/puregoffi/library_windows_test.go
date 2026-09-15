//go:build windows

package puregoffi

import "testing"

func TestLibraryLifecycle(t *testing.T) {
	handle, err := OpenLibrary("kernel32.dll")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSymbol(handle, "GetCurrentProcessId"); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSymbol(handle, "go_sdl3_symbol_that_does_not_exist"); err == nil {
		t.Fatal("missing symbol unexpectedly resolved")
	}
	if err := CloseLibrary(handle); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenLibrary("go-sdl3-library-that-does-not-exist.dll"); err == nil {
		t.Fatal("missing library unexpectedly opened")
	}
}
