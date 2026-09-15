// Package puregoffi provides the small runtime support layer needed by the
// checked-in SDL bindings.
package puregoffi

// BoolToUintptr converts a C-style boolean argument to its integer form.
func BoolToUintptr(value bool) uintptr {
	if value {
		return 1
	}
	return 0
}
