// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// PasswordEntryRowClass: instance of this type is always passed by reference.
type PasswordEntryRowClass struct {
	*passwordEntryRowClass
}

// passwordEntryRowClass is the struct that's finalized.
type passwordEntryRowClass struct {
	native *C.AdwPasswordEntryRowClass
}

func (p *PasswordEntryRowClass) ParentClass() *EntryRowClass {
	valptr := &p.native.parent_class
	var _v *EntryRowClass // out
	_v = (*EntryRowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
