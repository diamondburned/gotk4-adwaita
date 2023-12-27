// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
// extern void _gotk4_adw1_TabButton_ConnectClicked(gpointer, guintptr);
// extern void _gotk4_adw1_TabButton_ConnectActivate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeTabButton = coreglib.Type(C.adw_tab_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTabButton, F: marshalTabButton},
	})
}

// TabButtonOverrides contains methods that are overridable.
type TabButtonOverrides struct {
}

func defaultTabButtonOverrides(v *TabButton) TabButtonOverrides {
	return TabButtonOverrides{}
}

// TabButton: button that displays the number of tabview pages.
//
// <picture> <source srcset="tab-button-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="tab-button.png" alt="tab-button"> </picture>
//
// AdwTabButton is a button that displays the number of pages in a given
// AdwTabView, as well as whether one of the inactive pages needs attention.
//
// It's intended to be used as a visible indicator when there's no visible tab
// bar, typically opening an taboverview on click, e.g. via the overview.open
// action name:
//
//    <object class="AdwTabButton">
//      <property name="view">view</property>
//      <property name="action-name">overview.open</property>
//    </object>
//
// # CSS nodes
//
// AdwTabButton has a main CSS node with name tabbutton.
//
// # Accessibility
//
// AdwTabButton uses the GTK_ACCESSIBLE_ROLE_BUTTON role.
type TabButton struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Actionable
}

var (
	_ gtk.Widgetter     = (*TabButton)(nil)
	_ coreglib.Objector = (*TabButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TabButton, *TabButtonClass, TabButtonOverrides](
		GTypeTabButton,
		initTabButtonClass,
		wrapTabButton,
		defaultTabButtonOverrides,
	)
}

func initTabButtonClass(gclass unsafe.Pointer, overrides TabButtonOverrides, classInitFunc func(*TabButtonClass)) {
	if classInitFunc != nil {
		class := (*TabButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTabButton(obj *coreglib.Object) *TabButton {
	return &TabButton{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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
		Actionable: gtk.Actionable{
			Widget: gtk.Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
		},
	}
}

func marshalTabButton(p uintptr) (interface{}, error) {
	return wrapTabButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate is emitted to animate press then release.
//
// This is an action signal. Applications should never connect to this signal,
// but use the tabbutton::clicked signal.
func (self *TabButton) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "activate", false, unsafe.Pointer(C._gotk4_adw1_TabButton_ConnectActivate), f)
}

// ConnectClicked is emitted when the button has been activated (pressed and
// released).
func (self *TabButton) ConnectClicked(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "clicked", false, unsafe.Pointer(C._gotk4_adw1_TabButton_ConnectClicked), f)
}

// NewTabButton creates a new AdwTabButton.
//
// The function returns the following values:
//
//   - tabButton: newly created AdwTabButton.
//
func NewTabButton() *TabButton {
	var _cret *C.GtkWidget // in

	_cret = C.adw_tab_button_new()

	var _tabButton *TabButton // out

	_tabButton = wrapTabButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _tabButton
}

// View gets the tab view self displays.
//
// The function returns the following values:
//
//   - tabView (optional): tab view.
//
func (self *TabButton) View() *TabView {
	var _arg0 *C.AdwTabButton // out
	var _cret *C.AdwTabView   // in

	_arg0 = (*C.AdwTabButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_tab_button_get_view(_arg0)
	runtime.KeepAlive(self)

	var _tabView *TabView // out

	if _cret != nil {
		_tabView = wrapTabView(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _tabView
}

// SetView sets the tab view to display.
//
// The function takes the following parameters:
//
//   - view (optional): tab view.
//
func (self *TabButton) SetView(view *TabView) {
	var _arg0 *C.AdwTabButton // out
	var _arg1 *C.AdwTabView   // out

	_arg0 = (*C.AdwTabButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if view != nil {
		_arg1 = (*C.AdwTabView)(unsafe.Pointer(coreglib.InternObject(view).Native()))
	}

	C.adw_tab_button_set_view(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(view)
}