// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_status_page_get_type()), F: marshalStatusPager},
	})
}

// StatusPage: page used for empty/error states and similar use-cases.
//
// The AdwStatusPage widget can have an icon, a title, a description and a
// custom widget which is displayed below them.
//
//
// CSS nodes
//
// AdwStatusPage has a main CSS node with name statuspage.
//
// AdwStatusPage can use the .compact style class for when it needs to fit into
// a small space such a sidebar or a popover.
type StatusPage struct {
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*StatusPage)(nil)
)

func wrapStatusPage(obj *externglib.Object) *StatusPage {
	return &StatusPage{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalStatusPager(p uintptr) (interface{}, error) {
	return wrapStatusPage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStatusPage creates a new AdwStatusPage.
func NewStatusPage() *StatusPage {
	var _cret *C.GtkWidget // in

	_cret = C.adw_status_page_new()

	var _statusPage *StatusPage // out

	_statusPage = wrapStatusPage(externglib.Take(unsafe.Pointer(_cret)))

	return _statusPage
}

// Child gets the child widget of self.
func (self *StatusPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Description gets the description for self.
func (self *StatusPage) Description() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_description(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IconName gets the icon name for self.
func (self *StatusPage) IconName() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title for self.
func (self *StatusPage) Title() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child widget.
//
func (self *StatusPage) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.adw_status_page_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetDescription sets the description for self.
//
// The function takes the following parameters:
//
//    - description: description.
//
func (self *StatusPage) SetDescription(description string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	if description != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_status_page_set_description(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(description)
}

// SetIconName sets the icon name for self.
//
// The function takes the following parameters:
//
//    - iconName: icon name.
//
func (self *StatusPage) SetIconName(iconName string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_status_page_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetTitle sets the title for self.
//
// The function takes the following parameters:
//
//    - title: title.
//
func (self *StatusPage) SetTitle(title string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_status_page_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}
