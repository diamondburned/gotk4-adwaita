// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// glib.Type values for adw-application.go.
var GTypeApplication = externglib.Type(C.adw_application_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeApplication, F: marshalApplication},
	})
}

// ApplicationOverrider contains methods that are overridable.
type ApplicationOverrider interface {
}

// Application: base class for Adwaita applications.
//
// AdwApplication handles library initialization by calling init in the default
// gio.Application::startup signal handler, in turn chaining up as required by
// gtk.Application. Therefore, any subclass of AdwApplication should always
// chain up its startup handler before using any Adwaita or GTK API.
//
//
// Automatic Resources
//
// AdwApplication will automatically load stylesheets located in the
// application's resource base path (see gio.Application.SetResourceBasePath(),
// if they're present.
//
// They can be used to add custom styles to the application, as follows:
//
// - style.css contains styles that are always present.
//
// - style-dark.css contains styles only used when stylemanager:dark is TRUE.
//
// - style-hc.css contains styles used when the system high contrast preference
// is enabled.
//
// - style-hc-dark.css contains styles used when the system high contrast
// preference is enabled and stylemanager:dark is TRUE.
type Application struct {
	_ [0]func() // equal guard
	gtk.Application
}

var (
	_ externglib.Objector = (*Application)(nil)
)

func classInitApplicationer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapApplication(obj *externglib.Object) *Application {
	return &Application{
		Application: gtk.Application{
			Application: gio.Application{
				Object: obj,
				ActionGroup: gio.ActionGroup{
					Object: obj,
				},
				ActionMap: gio.ActionMap{
					Object: obj,
				},
			},
		},
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	return wrapApplication(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewApplication creates a new AdwApplication.
//
// If application_id is not NULL, then it must be valid. See
// gio.Application().IDIsValid.
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled.
//
// The function takes the following parameters:
//
//    - applicationId (optional): application ID.
//    - flags: application flags.
//
// The function returns the following values:
//
//    - application: newly created AdwApplication.
//
func NewApplication(applicationId string, flags gio.ApplicationFlags) *Application {
	var _arg1 *C.char             // out
	var _arg2 C.GApplicationFlags // out
	var _cret *C.AdwApplication   // in

	if applicationId != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(applicationId)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GApplicationFlags(flags)

	_cret = C.adw_application_new(_arg1, _arg2)
	runtime.KeepAlive(applicationId)
	runtime.KeepAlive(flags)

	var _application *Application // out

	_application = wrapApplication(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _application
}

// StyleManager gets the style manager for self.
//
// The function returns the following values:
//
//    - styleManager: style manager.
//
func (self *Application) StyleManager() *StyleManager {
	var _arg0 *C.AdwApplication  // out
	var _cret *C.AdwStyleManager // in

	_arg0 = (*C.AdwApplication)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_application_get_style_manager(_arg0)
	runtime.KeepAlive(self)

	var _styleManager *StyleManager // out

	_styleManager = wrapStyleManager(externglib.Take(unsafe.Pointer(_cret)))

	return _styleManager
}
