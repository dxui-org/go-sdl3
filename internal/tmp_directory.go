package internal

import (
	"os"
	"sync"
)

type tmpDir struct {
	onceCreate sync.Once
	onceRemove sync.Once

	Dir string
	Err error
}

var dir tmpDir

func TmpDir() (string, error) {
	dir.onceCreate.Do(func() {
		dir.Dir, dir.Err = os.MkdirTemp("", "")
	})

	return dir.Dir, dir.Err
}

func RemoveTmpDir() {
	dir.onceRemove.Do(func() {
		os.RemoveAll(dir.Dir)
	})
	// Clear tmpDir once entry after removal, so that it can be created again
	dir = tmpDir{}
}
