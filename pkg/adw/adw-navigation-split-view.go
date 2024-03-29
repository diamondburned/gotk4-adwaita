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

// NavigationSplitViewClass: instance of this type is always passed by
// reference.
type NavigationSplitViewClass struct {
	*navigationSplitViewClass
}

// navigationSplitViewClass is the struct that's finalized.
type navigationSplitViewClass struct {
	native *C.AdwNavigationSplitViewClass
}

func (n *NavigationSplitViewClass) ParentClass() *gtk.WidgetClass {
	valptr := &n.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
