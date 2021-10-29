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
		{T: externglib.Type(C.adw_view_switcher_title_get_type()), F: marshalViewSwitcherTitler},
	})
}

// ViewSwitcherTitle: view switcher title.
//
// A widget letting you switch between multiple views contained by a
// adw.ViewStack via an adw.ViewSwitcher.
//
// It is designed to be used as the title widget of a adw.HeaderBar, and will
// display the window's title when the window is too narrow to fit the view
// switcher e.g. on mobile phones, or if there are less than two views.
//
// You can conveniently bind the adw.ViewSwitcherBar:reveal property to
// adw.ViewSwitcherTitle:title-visible to automatically reveal the view switcher
// bar when the title label is displayed in place of the view switcher.
//
// An example of the UI definition for a common use case:
//
//    <object class="GtkWindow"/>
//      <child type="titlebar">
//        <object class="AdwHeaderBar">
//          <property name="centering-policy">strict</property>
//          <child type="title">
//            <object class="AdwViewSwitcherTitle" id="title">
//              <property name="stack">stack</property>
//            </object>
//          </child>
//        </object>
//      </child>
//      <child>
//        <object class="GtkBox">
//          <child>
//            <object class="AdwViewStack" id="stack"/>
//          </child>
//          <child>
//            <object class="AdwViewSwitcherBar">
//              <property name="stack">stack</property>
//              <binding name="reveal">
//                <lookup name="title-visible">title</lookup>
//              </binding>
//            </object>
//          </child>
//        </object>
//      </child>
//    </object>
//
//
//
// CSS nodes
//
// AdwViewSwitcherTitle has a single CSS node with name viewswitchertitle.
type ViewSwitcherTitle struct {
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*ViewSwitcherTitle)(nil)
)

func wrapViewSwitcherTitle(obj *externglib.Object) *ViewSwitcherTitle {
	return &ViewSwitcherTitle{
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

func marshalViewSwitcherTitler(p uintptr) (interface{}, error) {
	return wrapViewSwitcherTitle(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewSwitcherTitle creates a new AdwViewSwitcherTitle.
func NewViewSwitcherTitle() *ViewSwitcherTitle {
	var _cret *C.GtkWidget // in

	_cret = C.adw_view_switcher_title_new()

	var _viewSwitcherTitle *ViewSwitcherTitle // out

	_viewSwitcherTitle = wrapViewSwitcherTitle(externglib.Take(unsafe.Pointer(_cret)))

	return _viewSwitcherTitle
}

// Stack gets the stack controlled by self.
func (self *ViewSwitcherTitle) Stack() *ViewStack {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _cret *C.AdwViewStack         // in

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_title_get_stack(_arg0)
	runtime.KeepAlive(self)

	var _viewStack *ViewStack // out

	if _cret != nil {
		_viewStack = wrapViewStack(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _viewStack
}

// Subtitle gets the subtitle of self.
func (self *ViewSwitcherTitle) Subtitle() string {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _cret *C.char                 // in

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_title_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Title gets the title of self.
func (self *ViewSwitcherTitle) Title() string {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _cret *C.char                 // in

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_title_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TitleVisible gets whether the title of self is currently visible.
func (self *ViewSwitcherTitle) TitleVisible() bool {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _cret C.gboolean              // in

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_title_get_title_visible(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ViewSwitcherEnabled gets whether self's view switcher is enabled.
func (self *ViewSwitcherTitle) ViewSwitcherEnabled() bool {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _cret C.gboolean              // in

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_title_get_view_switcher_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetStack sets the stack controlled by self.
//
// The function takes the following parameters:
//
//    - stack: stack.
//
func (self *ViewSwitcherTitle) SetStack(stack *ViewStack) {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _arg1 *C.AdwViewStack         // out

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	if stack != nil {
		_arg1 = (*C.AdwViewStack)(unsafe.Pointer(stack.Native()))
	}

	C.adw_view_switcher_title_set_stack(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(stack)
}

// SetSubtitle sets the subtitle of self.
//
// The function takes the following parameters:
//
//    - subtitle: subtitle.
//
func (self *ViewSwitcherTitle) SetSubtitle(subtitle string) {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_view_switcher_title_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of self.
//
// The function takes the following parameters:
//
//    - title: title.
//
func (self *ViewSwitcherTitle) SetTitle(title string) {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_view_switcher_title_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetViewSwitcherEnabled sets whether self's view switcher is enabled.
//
// The function takes the following parameters:
//
//    - enabled: whether the view switcher is enabled.
//
func (self *ViewSwitcherTitle) SetViewSwitcherEnabled(enabled bool) {
	var _arg0 *C.AdwViewSwitcherTitle // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.AdwViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.adw_view_switcher_title_set_view_switcher_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enabled)
}
