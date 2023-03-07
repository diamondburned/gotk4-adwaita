// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
// extern void _gotk4_adw1_EntryRow_ConnectEntryActivated(gpointer, guintptr);
// extern void _gotk4_adw1_EntryRow_ConnectApply(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeEntryRow = coreglib.Type(C.adw_entry_row_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEntryRow, F: marshalEntryRow},
	})
}

// EntryRowOverrides contains methods that are overridable.
type EntryRowOverrides struct {
}

func defaultEntryRowOverrides(v *EntryRow) EntryRowOverrides {
	return EntryRowOverrides{}
}

// EntryRow: gtk.ListBoxRow with an embedded text entry.
//
// <picture> <source srcset="entry-row-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="entry-row.png" alt="entry-row"> </picture>
//
// AdwEntryRow has a title that doubles as placeholder text. It shows an icon
// indicating that it's editable and can receive additional widgets before or
// after the editable part.
//
// If entryrow:show-apply-button is set to TRUE, AdwEntryRow can show an apply
// button when editing its contents. This can be useful if changing its contents
// can result in an expensive operation, such as network activity.
//
// AdwEntryRow provides only minimal API and should be used with the
// gtk.Editable API.
//
// See also passwordentryrow.
//
// # AdwEntryRow as GtkBuildable
//
// The AdwEntryRow implementation of the gtk.Buildable interface supports adding
// a child at its end by specifying “suffix” or omitting the “type” attribute of
// a <child> element.
//
// It also supports adding a child as a prefix widget by specifying “prefix” as
// the “type” attribute of a <child> element.
//
// # CSS nodes
//
// AdwEntryRow has a single CSS node with name row and the .entry style class.
type EntryRow struct {
	_ [0]func() // equal guard
	PreferencesRow

	*coreglib.Object
	gtk.Editable
}

var (
	_ coreglib.Objector = (*EntryRow)(nil)
	_ gtk.Widgetter     = (*EntryRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EntryRow, *EntryRowClass, EntryRowOverrides](
		GTypeEntryRow,
		initEntryRowClass,
		wrapEntryRow,
		defaultEntryRowOverrides,
	)
}

func initEntryRowClass(gclass unsafe.Pointer, overrides EntryRowOverrides, classInitFunc func(*EntryRowClass)) {
	if classInitFunc != nil {
		class := (*EntryRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEntryRow(obj *coreglib.Object) *EntryRow {
	return &EntryRow{
		PreferencesRow: PreferencesRow{
			ListBoxRow: gtk.ListBoxRow{
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
			},
		},
		Object: obj,
		Editable: gtk.Editable{
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

func marshalEntryRow(p uintptr) (interface{}, error) {
	return wrapEntryRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectApply is emitted when the apply button is pressed.
//
// See entryrow:show-apply-button.
func (self *EntryRow) ConnectApply(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "apply", false, unsafe.Pointer(C._gotk4_adw1_EntryRow_ConnectApply), f)
}

// ConnectEntryActivated is emitted when the embedded entry is activated.
func (self *EntryRow) ConnectEntryActivated(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "entry-activated", false, unsafe.Pointer(C._gotk4_adw1_EntryRow_ConnectEntryActivated), f)
}

// NewEntryRow creates a new AdwEntryRow.
//
// The function returns the following values:
//
//   - entryRow: newly created AdwEntryRow.
//
func NewEntryRow() *EntryRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_entry_row_new()

	var _entryRow *EntryRow // out

	_entryRow = wrapEntryRow(coreglib.Take(unsafe.Pointer(_cret)))

	return _entryRow
}

// AddPrefix adds a prefix widget to self.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *EntryRow) AddPrefix(widget gtk.Widgetter) {
	var _arg0 *C.AdwEntryRow // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_entry_row_add_prefix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// AddSuffix adds a suffix widget to self.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *EntryRow) AddSuffix(widget gtk.Widgetter) {
	var _arg0 *C.AdwEntryRow // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_entry_row_add_suffix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// ActivatesDefault gets whether activating the embedded entry can activate the
// default widget.
//
// The function returns the following values:
//
//   - ok: whether to activate the default widget.
//
func (self *EntryRow) ActivatesDefault() bool {
	var _arg0 *C.AdwEntryRow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_entry_row_get_activates_default(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Attributes gets Pango attributes applied to the text of the embedded entry.
//
// The function returns the following values:
//
//   - attrList (optional): list of attributes.
//
func (self *EntryRow) Attributes() *pango.AttrList {
	var _arg0 *C.AdwEntryRow   // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_entry_row_get_attributes(_arg0)
	runtime.KeepAlive(self)

	var _attrList *pango.AttrList // out

	if _cret != nil {
		_attrList = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_attrList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
			},
		)
	}

	return _attrList
}

// EnableEmojiCompletion gets whether to suggest emoji replacements on self.
//
// The function returns the following values:
//
//   - ok: whether or not emoji completion is enabled.
//
func (self *EntryRow) EnableEmojiCompletion() bool {
	var _arg0 *C.AdwEntryRow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_entry_row_get_enable_emoji_completion(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InputHints gets the additional input hints of self.
//
// The function returns the following values:
//
//   - inputHints: input hints.
//
func (self *EntryRow) InputHints() gtk.InputHints {
	var _arg0 *C.AdwEntryRow  // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_entry_row_get_input_hints(_arg0)
	runtime.KeepAlive(self)

	var _inputHints gtk.InputHints // out

	_inputHints = gtk.InputHints(_cret)

	return _inputHints
}

// InputPurpose gets the input purpose of self.
//
// The function returns the following values:
//
//   - inputPurpose: input purpose.
//
func (self *EntryRow) InputPurpose() gtk.InputPurpose {
	var _arg0 *C.AdwEntryRow    // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_entry_row_get_input_purpose(_arg0)
	runtime.KeepAlive(self)

	var _inputPurpose gtk.InputPurpose // out

	_inputPurpose = gtk.InputPurpose(_cret)

	return _inputPurpose
}

// ShowApplyButton gets whether self can show the apply button.
//
// The function returns the following values:
//
//   - ok: whether to show the apply button.
//
func (self *EntryRow) ShowApplyButton() bool {
	var _arg0 *C.AdwEntryRow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_entry_row_get_show_apply_button(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes a child from self.
//
// The function takes the following parameters:
//
//   - widget: child to be removed.
//
func (self *EntryRow) Remove(widget gtk.Widgetter) {
	var _arg0 *C.AdwEntryRow // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_entry_row_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetActivatesDefault sets whether activating the embedded entry can activate
// the default widget.
//
// The function takes the following parameters:
//
//   - activates: whether to activate the default widget.
//
func (self *EntryRow) SetActivatesDefault(activates bool) {
	var _arg0 *C.AdwEntryRow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if activates {
		_arg1 = C.TRUE
	}

	C.adw_entry_row_set_activates_default(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(activates)
}

// SetAttributes sets Pango attributes to apply to the text of the embedded
// entry.
//
// The pango.Attribute's start_index and end_index must refer to the
// gtk.EntryBuffer text, i.e. without the preedit string.
//
// The function takes the following parameters:
//
//   - attributes (optional): list of attributes.
//
func (self *EntryRow) SetAttributes(attributes *pango.AttrList) {
	var _arg0 *C.AdwEntryRow   // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if attributes != nil {
		_arg1 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attributes)))
	}

	C.adw_entry_row_set_attributes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(attributes)
}

// SetEnableEmojiCompletion sets whether to suggest emoji replacements on self.
//
// Emoji replacement is done with :-delimited names, like :heart:.
//
// The function takes the following parameters:
//
//   - enableEmojiCompletion: whether emoji completion should be enabled or not.
//
func (self *EntryRow) SetEnableEmojiCompletion(enableEmojiCompletion bool) {
	var _arg0 *C.AdwEntryRow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enableEmojiCompletion {
		_arg1 = C.TRUE
	}

	C.adw_entry_row_set_enable_emoji_completion(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableEmojiCompletion)
}

// SetInputHints: set additional input hints for self.
//
// Input hints allow input methods to fine-tune their behavior.
//
// See also: adwentryrow:input-purpose.
//
// The function takes the following parameters:
//
//   - hints: hints.
//
func (self *EntryRow) SetInputHints(hints gtk.InputHints) {
	var _arg0 *C.AdwEntryRow  // out
	var _arg1 C.GtkInputHints // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkInputHints(hints)

	C.adw_entry_row_set_input_hints(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(hints)
}

// SetInputPurpose sets the input purpose of self.
//
// The input purpose can be used by input methods to adjust their behavior.
//
// The function takes the following parameters:
//
//   - purpose: purpose.
//
func (self *EntryRow) SetInputPurpose(purpose gtk.InputPurpose) {
	var _arg0 *C.AdwEntryRow    // out
	var _arg1 C.GtkInputPurpose // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkInputPurpose(purpose)

	C.adw_entry_row_set_input_purpose(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(purpose)
}

// SetShowApplyButton sets whether self can show the apply button.
//
// When set to TRUE, typing text in the entry will reveal an apply button.
// Clicking it or pressing the <kbd>Enter</kbd> key will hide the button and
// emit the entryrow::apply signal.
//
// This is useful if changing the entry contents can trigger an expensive
// operation, e.g. network activity, to avoid triggering it after typing every
// character.
//
// The function takes the following parameters:
//
//   - showApplyButton: whether to show the apply button.
//
func (self *EntryRow) SetShowApplyButton(showApplyButton bool) {
	var _arg0 *C.AdwEntryRow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwEntryRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if showApplyButton {
		_arg1 = C.TRUE
	}

	C.adw_entry_row_set_show_apply_button(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showApplyButton)
}
