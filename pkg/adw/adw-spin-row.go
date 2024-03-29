// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// SpinRowClass: instance of this type is always passed by reference.
type SpinRowClass struct {
	*spinRowClass
}

// spinRowClass is the struct that's finalized.
type spinRowClass struct {
	native *C.AdwSpinRowClass
}

func (s *SpinRowClass) ParentClass() *ActionRowClass {
	valptr := &s.native.parent_class
	var _v *ActionRowClass // out
	_v = (*ActionRowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
