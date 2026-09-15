//go:build windows

package puregoffi

import "golang.org/x/sys/windows"

// BytePtrFromString returns a pointer to a newly allocated, NUL-terminated
// copy of value. The allocation remains live while the returned pointer is
// reachable. It panics when value contains an embedded NUL byte.
func BytePtrFromString(value string) *byte {
	ptr, err := windows.BytePtrFromString(value)
	if err != nil {
		panic(err)
	}
	return ptr
}

// BytePtrToString copies a NUL-terminated byte sequence into a Go string.
// A nil pointer produces the empty string.
func BytePtrToString(ptr *byte) string {
	return windows.BytePtrToString(ptr)
}
