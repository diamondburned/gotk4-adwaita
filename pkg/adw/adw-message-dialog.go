// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// MessageDialogClass: instance of this type is always passed by reference.
type MessageDialogClass struct {
	*messageDialogClass
}

// messageDialogClass is the struct that's finalized.
type messageDialogClass struct {
	native *C.AdwMessageDialogClass
}

func (m *MessageDialogClass) ParentClass() *gtk.WindowClass {
	valptr := &m.native.parent_class
	var _v *gtk.WindowClass // out
	_v = (*gtk.WindowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
