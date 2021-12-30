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
		{T: externglib.Type(C.adw_window_title_get_type()), F: marshalWindowTitler},
	})
}

// WindowTitle: helper widget for setting a window's title and subtitle.
//
// AdwWindowTitle shows a title and subtitle. It's intended to be used as the
// title child of gtk.HeaderBar or adw.HeaderBar.
//
//
// CSS nodes
//
// AdwWindowTitle has a single CSS node with name windowtitle.
type WindowTitle struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*WindowTitle)(nil)
)

func wrapWindowTitle(obj *externglib.Object) *WindowTitle {
	return &WindowTitle{
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
	}
}

func marshalWindowTitler(p uintptr) (interface{}, error) {
	return wrapWindowTitle(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowTitle creates a new AdwWindowTitle.
//
// The function takes the following parameters:
//
//    - title: title.
//    - subtitle: subtitle.
//
// The function returns the following values:
//
//    - windowTitle: newly created AdwWindowTitle.
//
func NewWindowTitle(title, subtitle string) *WindowTitle {
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.adw_window_title_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(subtitle)

	var _windowTitle *WindowTitle // out

	_windowTitle = wrapWindowTitle(externglib.Take(unsafe.Pointer(_cret)))

	return _windowTitle
}

// Subtitle gets the subtitle of self.
//
// The function returns the following values:
//
//    - utf8: subtitle.
//
func (self *WindowTitle) Subtitle() string {
	var _arg0 *C.AdwWindowTitle // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_window_title_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Title gets the title of self.
//
// The function returns the following values:
//
//    - utf8: title.
//
func (self *WindowTitle) Title() string {
	var _arg0 *C.AdwWindowTitle // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_window_title_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetSubtitle sets the subtitle of self.
//
// The function takes the following parameters:
//
//    - subtitle: subtitle.
//
func (self *WindowTitle) SetSubtitle(subtitle string) {
	var _arg0 *C.AdwWindowTitle // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_window_title_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of self.
//
// The function takes the following parameters:
//
//    - title: title.
//
func (self *WindowTitle) SetTitle(title string) {
	var _arg0 *C.AdwWindowTitle // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_window_title_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}
