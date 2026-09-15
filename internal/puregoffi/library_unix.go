//go:build linux || darwin

package puregoffi

import "github.com/ebitengine/purego"

func OpenLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}

func OpenSymbol(library uintptr, name string) (uintptr, error) {
	return purego.Dlsym(library, name)
}

func CloseLibrary(library uintptr) error {
	return purego.Dlclose(library)
}
