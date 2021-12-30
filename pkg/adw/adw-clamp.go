// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_clamp_get_type()), F: marshalClamper},
	})
}

// Clamp: widget constraining its child to a given size.
//
// The AdwClamp widget constrains the size of the widget it contains to a given
// maximum size. It will constrain the width if it is horizontal, or the height
// if it is vertical. The expansion of the child from its minimum to its maximum
// size is eased out for a smooth transition.
//
// If the child requires more than the requested maximum size, it will be
// allocated the minimum size it can fit in instead.
//
//
// CSS nodes
//
// AdwClamp has a single CSS node with name clamp.
//
// Its children will receive the style classes .large when the child reached its
// maximum size, .small when the clamp allocates its full size to the child,
// .medium in-between, or none if it hasn't computed its size yet.
type Clamp struct {
	_ [0]func() // equal guard
	gtk.Widget

	*externglib.Object
	gtk.Orientable
}

var (
	_ gtk.Widgetter       = (*Clamp)(nil)
	_ externglib.Objector = (*Clamp)(nil)
)

func wrapClamp(obj *externglib.Object) *Clamp {
	return &Clamp{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalClamper(p uintptr) (interface{}, error) {
	return wrapClamp(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewClamp creates a new AdwClamp.
//
// The function returns the following values:
//
//    - clamp: newly created AdwClamp.
//
func NewClamp() *Clamp {
	var _cret *C.GtkWidget // in

	_cret = C.adw_clamp_new()

	var _clamp *Clamp // out

	_clamp = wrapClamp(externglib.Take(unsafe.Pointer(_cret)))

	return _clamp
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self.
//
func (self *Clamp) Child() gtk.Widgetter {
	var _arg0 *C.AdwClamp  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(self.Native()))

	_cret = C.adw_clamp_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// MaximumSize gets the maximum size allocated to the child.
//
// The function returns the following values:
//
//    - gint: maximum size to allocate to the child.
//
func (self *Clamp) MaximumSize() int {
	var _arg0 *C.AdwClamp // out
	var _cret C.int       // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(self.Native()))

	_cret = C.adw_clamp_get_maximum_size(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TighteningThreshold gets the size above which the child is clamped.
//
// The function returns the following values:
//
//    - gint: size above which the child is clamped.
//
func (self *Clamp) TighteningThreshold() int {
	var _arg0 *C.AdwClamp // out
	var _cret C.int       // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(self.Native()))

	_cret = C.adw_clamp_get_tightening_threshold(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (self *Clamp) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwClamp  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.adw_clamp_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetMaximumSize sets the maximum size allocated to the child.
//
// The function takes the following parameters:
//
//    - maximumSize: maximum size.
//
func (self *Clamp) SetMaximumSize(maximumSize int) {
	var _arg0 *C.AdwClamp // out
	var _arg1 C.int       // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(maximumSize)

	C.adw_clamp_set_maximum_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maximumSize)
}

// SetTighteningThreshold sets the size above which the child is clamped.
//
// The function takes the following parameters:
//
//    - tighteningThreshold: tightening threshold.
//
func (self *Clamp) SetTighteningThreshold(tighteningThreshold int) {
	var _arg0 *C.AdwClamp // out
	var _arg1 C.int       // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(tighteningThreshold)

	C.adw_clamp_set_tightening_threshold(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(tighteningThreshold)
}
