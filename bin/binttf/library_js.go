//go:build js

package binttf

import "github.com/dxui-org/go-sdl3/ttf"

type library struct{}

func Load() library {
	return library{}
}

func (l library) Unload() {
	ttf.CloseLibrary()
}
