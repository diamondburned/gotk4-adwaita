// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.adw_application_window_get_type()), F: marshalApplicationWindower},
	})
}

// ApplicationWindow: freeform application window.
//
// AdwApplicationWindow is a gtk.ApplicationWindow subclass providing the same
// features as adw.Window.
//
// See adw.Window for details.
//
// Using gtk.Application:menubar is not supported and may result in visual
// glitches.
type ApplicationWindow struct {
	gtk.ApplicationWindow
}

var (
	_ externglib.Objector = (*ApplicationWindow)(nil)
	_ gtk.Widgetter       = (*ApplicationWindow)(nil)
)

func wrapApplicationWindow(obj *externglib.Object) *ApplicationWindow {
	return &ApplicationWindow{
		ApplicationWindow: gtk.ApplicationWindow{
			Window: gtk.Window{
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
				Root: gtk.Root{
					NativeSurface: gtk.NativeSurface{
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
					},
				},
				ShortcutManager: gtk.ShortcutManager{
					Object: obj,
				},
				Object: obj,
			},
			ActionGroup: gio.ActionGroup{
				Object: obj,
			},
			ActionMap: gio.ActionMap{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalApplicationWindower(p uintptr) (interface{}, error) {
	return wrapApplicationWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewApplicationWindow creates a new AdwApplicationWindow for app.
//
// The function takes the following parameters:
//
//    - app: application instance.
//
func NewApplicationWindow(app *gtk.Application) *ApplicationWindow {
	var _arg1 *C.GtkApplication // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkApplication)(unsafe.Pointer(app.Native()))

	_cret = C.adw_application_window_new(_arg1)
	runtime.KeepAlive(app)

	var _applicationWindow *ApplicationWindow // out

	_applicationWindow = wrapApplicationWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _applicationWindow
}

// Content gets the content widget of self.
//
// This method should always be used instead of gtk.Window.GetChild().
func (self *ApplicationWindow) Content() gtk.Widgetter {
	var _arg0 *C.AdwApplicationWindow // out
	var _cret *C.GtkWidget            // in

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_application_window_get_content(_arg0)
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

// SetContent sets the content widget of self.
//
// This method should always be used instead of gtk.Window.SetChild().
//
// The function takes the following parameters:
//
//    - content widget.
//
func (self *ApplicationWindow) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwApplicationWindow // out
	var _arg1 *C.GtkWidget            // out

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(self.Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(content.Native()))
	}

	C.adw_application_window_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}
