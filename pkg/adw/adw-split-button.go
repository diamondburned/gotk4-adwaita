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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_split_button_get_type()), F: marshalSplitButtonner},
	})
}

// SplitButton: combined button and dropdown widget.
//
// AdwSplitButton is typically used to present a set of actions in a menu, but
// allow access to one of them with a single click.
//
// The API is very similar to gtk.Button and gtk.MenuButton, see their
// documentation for details.
//
// CSS nodes
//
//    splitbutton[.image-button][.text-button]
//    ├── button
//    │   ╰── <content>
//    ├── separator
//    ╰── menubutton
//        ╰── button.toggle
//            ╰── arrow
//
//
// AdwSplitButton's CSS node is called splitbutton. It contains the css nodes:
// button, separator, menubutton. See gtk.MenuButton documentation for the
// menubutton contents.
//
// The main CSS node will contain the .image-button or .text-button style
// classes matching the button contents. The nested button nodes will never
// contain them.
//
//
// Accessibility
//
// AdwSplitButton uses the GTK_ACCESSIBLE_ROLE_BUTTON role.
type SplitButton struct {
	_ [0]func() // equal guard
	gtk.Widget

	*externglib.Object
	gtk.Actionable
}

var (
	_ gtk.Widgetter       = (*SplitButton)(nil)
	_ externglib.Objector = (*SplitButton)(nil)
)

func wrapSplitButton(obj *externglib.Object) *SplitButton {
	return &SplitButton{
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
		Actionable: gtk.Actionable{
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
		},
	}
}

func marshalSplitButtonner(p uintptr) (interface{}, error) {
	return wrapSplitButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate: emitted to animate press then release.
//
// This is an action signal. Applications should never connect to this signal,
// but use the adw.SplitButton::clicked signal.
func (self *SplitButton) ConnectActivate(f func()) externglib.SignalHandle {
	return self.Connect("activate", f)
}

// ConnectClicked: emitted when the button has been activated (pressed and
// released).
func (self *SplitButton) ConnectClicked(f func()) externglib.SignalHandle {
	return self.Connect("clicked", f)
}

// NewSplitButton creates a new AdwSplitButton.
//
// The function returns the following values:
//
//    - splitButton: newly created AdwSplitButton.
//
func NewSplitButton() *SplitButton {
	var _cret *C.GtkWidget // in

	_cret = C.adw_split_button_new()

	var _splitButton *SplitButton // out

	_splitButton = wrapSplitButton(externglib.Take(unsafe.Pointer(_cret)))

	return _splitButton
}

// Child gets the child widget.
//
// The function returns the following values:
//
//    - widget (optional): child widget.
//
func (self *SplitButton) Child() gtk.Widgetter {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_child(_arg0)
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

// Direction gets the direction in which the popup will be popped up.
//
// The function returns the following values:
//
//    - arrowType: direction.
//
func (self *SplitButton) Direction() gtk.ArrowType {
	var _arg0 *C.AdwSplitButton // out
	var _cret C.GtkArrowType    // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_direction(_arg0)
	runtime.KeepAlive(self)

	var _arrowType gtk.ArrowType // out

	_arrowType = gtk.ArrowType(_cret)

	return _arrowType
}

// IconName gets the name of the icon used to automatically populate the button.
//
// If the icon name has not been set with adw.SplitButton.SetIconName() the
// return value will be NULL.
//
// The function returns the following values:
//
//    - utf8 (optional): icon name.
//
func (self *SplitButton) IconName() string {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Label gets the label for self.
//
// The function returns the following values:
//
//    - utf8 (optional): label for self.
//
func (self *SplitButton) Label() string {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// MenuModel gets the menu model from which the popup will be created.
//
// The function returns the following values:
//
//    - menuModel (optional): menu model.
//
func (self *SplitButton) MenuModel() gio.MenuModeller {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_menu_model(_arg0)
	runtime.KeepAlive(self)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// Popover gets the popover that will be popped up when the dropdown is clicked.
//
// The function returns the following values:
//
//    - popover (optional): popover.
//
func (self *SplitButton) Popover() *gtk.Popover {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.GtkPopover     // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_popover(_arg0)
	runtime.KeepAlive(self)

	var _popover *gtk.Popover // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_popover = &gtk.Popover{
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
				NativeSurface: gtk.NativeSurface{
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
				},
				ShortcutManager: gtk.ShortcutManager{
					Object: obj,
				},
			}
		}
	}

	return _popover
}

// UseUnderline gets whether an underline in the text indicates a mnemonic.
//
// The function returns the following values:
//
//    - ok: whether an underline in the text indicates a mnemonic.
//
func (self *SplitButton) UseUnderline() bool {
	var _arg0 *C.AdwSplitButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	_cret = C.adw_split_button_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown dismisses the menu.
func (self *SplitButton) Popdown() {
	var _arg0 *C.AdwSplitButton // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	C.adw_split_button_popdown(_arg0)
	runtime.KeepAlive(self)
}

// Popup pops up the menu.
func (self *SplitButton) Popup() {
	var _arg0 *C.AdwSplitButton // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))

	C.adw_split_button_popup(_arg0)
	runtime.KeepAlive(self)
}

// SetChild sets the child widget.
//
// The function takes the following parameters:
//
//    - child (optional): new child widget.
//
func (self *SplitButton) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.adw_split_button_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetDirection sets the direction in which the popup will be popped up.
//
// The function takes the following parameters:
//
//    - direction: direction.
//
func (self *SplitButton) SetDirection(direction gtk.ArrowType) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 C.GtkArrowType    // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkArrowType(direction)

	C.adw_split_button_set_direction(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(direction)
}

// SetIconName sets the name of the icon used to automatically populate the
// button.
//
// The function takes the following parameters:
//
//    - iconName: icon name to set.
//
func (self *SplitButton) SetIconName(iconName string) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_split_button_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetLabel sets the label for self.
//
// The function takes the following parameters:
//
//    - label to set.
//
func (self *SplitButton) SetLabel(label string) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_split_button_set_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(label)
}

// SetMenuModel sets the menu model from which the popup will be created.
//
// The function takes the following parameters:
//
//    - menuModel (optional): menu model.
//
func (self *SplitButton) SetMenuModel(menuModel gio.MenuModeller) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	if menuModel != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(menuModel.Native()))
	}

	C.adw_split_button_set_menu_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(menuModel)
}

// SetPopover sets the popover that will be popped up when the dropdown is
// clicked.
//
// The function takes the following parameters:
//
//    - popover (optional): popover.
//
func (self *SplitButton) SetPopover(popover *gtk.Popover) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.GtkPopover     // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	if popover != nil {
		_arg1 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	}

	C.adw_split_button_set_popover(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(popover)
}

// SetUseUnderline sets whether an underline in the text indicates a mnemonic.
//
// The function takes the following parameters:
//
//    - useUnderline: whether an underline in the text indicates a mnemonic.
//
func (self *SplitButton) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_split_button_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}
