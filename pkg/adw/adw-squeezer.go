// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
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
		{T: externglib.Type(C.adw_squeezer_transition_type_get_type()), F: marshalSqueezerTransitionType},
		{T: externglib.Type(C.adw_squeezer_get_type()), F: marshalSqueezerer},
		{T: externglib.Type(C.adw_squeezer_page_get_type()), F: marshalSqueezerPager},
	})
}

// SqueezerTransitionType describes the possible transitions in a adw.Squeezer
// widget.
type SqueezerTransitionType C.gint

const (
	// SqueezerTransitionTypeNone: no transition.
	SqueezerTransitionTypeNone SqueezerTransitionType = iota
	// SqueezerTransitionTypeCrossfade: cross-fade.
	SqueezerTransitionTypeCrossfade
)

func marshalSqueezerTransitionType(p uintptr) (interface{}, error) {
	return SqueezerTransitionType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SqueezerTransitionType.
func (s SqueezerTransitionType) String() string {
	switch s {
	case SqueezerTransitionTypeNone:
		return "None"
	case SqueezerTransitionTypeCrossfade:
		return "Crossfade"
	default:
		return fmt.Sprintf("SqueezerTransitionType(%d)", s)
	}
}

// Squeezer: best fit container.
//
// The AdwSqueezer widget is a container which only shows the first of its
// children that fits in the available size. It is convenient to offer different
// widgets to represent the same data with different levels of detail, making
// the widget seem to squeeze itself to fit in the available space.
//
// Transitions between children can be animated as fades. This can be controlled
// with adw.Squeezer:transition-type.
//
//
// CSS nodes
//
// AdwSqueezer has a single CSS node with name squeezer.
type Squeezer struct {
	gtk.Widget

	gtk.Orientable
	*externglib.Object
}

var (
	_ gtk.Widgetter       = (*Squeezer)(nil)
	_ externglib.Objector = (*Squeezer)(nil)
)

func wrapSqueezer(obj *externglib.Object) *Squeezer {
	return &Squeezer{
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
		Orientable: gtk.Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSqueezerer(p uintptr) (interface{}, error) {
	return wrapSqueezer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSqueezer creates a new AdwSqueezer.
func NewSqueezer() *Squeezer {
	var _cret *C.GtkWidget // in

	_cret = C.adw_squeezer_new()

	var _squeezer *Squeezer // out

	_squeezer = wrapSqueezer(externglib.Take(unsafe.Pointer(_cret)))

	return _squeezer
}

// Add adds a child to self.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
func (self *Squeezer) Add(child gtk.Widgetter) *SqueezerPage {
	var _arg0 *C.AdwSqueezer     // out
	var _arg1 *C.GtkWidget       // out
	var _cret *C.AdwSqueezerPage // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_squeezer_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _squeezerPage *SqueezerPage // out

	_squeezerPage = wrapSqueezerPage(externglib.Take(unsafe.Pointer(_cret)))

	return _squeezerPage
}

// AllowNone gets whether to allow squeezing beyond the last child's minimum
// size.
func (self *Squeezer) AllowNone() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_allow_none(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether all children have the same size for the opposite
// orientation.
func (self *Squeezer) Homogeneous() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_homogeneous(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize gets whether self interpolates its size when changing the
// visible child.
func (self *Squeezer) InterpolateSize() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_interpolate_size(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Page returns the adw.SqueezerPage object for child.
//
// The function takes the following parameters:
//
//    - child of self.
//
func (self *Squeezer) Page(child gtk.Widgetter) *SqueezerPage {
	var _arg0 *C.AdwSqueezer     // out
	var _arg1 *C.GtkWidget       // out
	var _cret *C.AdwSqueezerPage // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_squeezer_get_page(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _squeezerPage *SqueezerPage // out

	_squeezerPage = wrapSqueezerPage(externglib.Take(unsafe.Pointer(_cret)))

	return _squeezerPage
}

// Pages returns a gio.ListModel that contains the pages of self.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track the visible page.
func (self *Squeezer) Pages() gtk.SelectionModeller {
	var _arg0 *C.AdwSqueezer       // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_pages(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel gtk.SelectionModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.SelectionModeller is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(gtk.SelectionModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.SelectionModeller")
		}
		_selectionModel = rv
	}

	return _selectionModel
}

// SwitchThresholdPolicy gets the fold threshold policy for self.
func (self *Squeezer) SwitchThresholdPolicy() FoldThresholdPolicy {
	var _arg0 *C.AdwSqueezer           // out
	var _cret C.AdwFoldThresholdPolicy // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_switch_threshold_policy(_arg0)
	runtime.KeepAlive(self)

	var _foldThresholdPolicy FoldThresholdPolicy // out

	_foldThresholdPolicy = FoldThresholdPolicy(_cret)

	return _foldThresholdPolicy
}

// TransitionDuration gets the transition animation duration for self.
func (self *Squeezer) TransitionDuration() uint {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.guint        // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_transition_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionRunning gets whether a transition is currently running for self.
func (self *Squeezer) TransitionRunning() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_transition_running(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionType gets the type of animation used for transitions between
// children in self.
func (self *Squeezer) TransitionType() SqueezerTransitionType {
	var _arg0 *C.AdwSqueezer              // out
	var _cret C.AdwSqueezerTransitionType // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_transition_type(_arg0)
	runtime.KeepAlive(self)

	var _squeezerTransitionType SqueezerTransitionType // out

	_squeezerTransitionType = SqueezerTransitionType(_cret)

	return _squeezerTransitionType
}

// VisibleChild gets the currently visible child of self.
func (self *Squeezer) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwSqueezer // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_visible_child(_arg0)
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

// XAlign gets the horizontal alignment, from 0 (start) to 1 (end).
func (self *Squeezer) XAlign() float32 {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.float        // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_xalign(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// YAlign gets the vertical alignment, from 0 (top) to 1 (bottom).
func (self *Squeezer) YAlign() float32 {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.float        // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_get_yalign(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Remove removes a child widget from self.
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (self *Squeezer) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_squeezer_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetAllowNone sets whether to allow squeezing beyond the last child's minimum
// size.
//
// The function takes the following parameters:
//
//    - allowNone: whether self allows squeezing beyond the last child.
//
func (self *Squeezer) SetAllowNone(allowNone bool) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	if allowNone {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_set_allow_none(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowNone)
}

// SetHomogeneous sets whether all children have the same size for the opposite
// orientation.
//
// The function takes the following parameters:
//
//    - homogeneous: whether self is homogeneous.
//
func (self *Squeezer) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(homogeneous)
}

// SetInterpolateSize sets whether self interpolates its size when changing the
// visible child.
//
// The function takes the following parameters:
//
//    - interpolateSize: whether to interpolate the size.
//
func (self *Squeezer) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_set_interpolate_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(interpolateSize)
}

// SetSwitchThresholdPolicy sets the fold threshold policy for self.
//
// The function takes the following parameters:
//
//    - policy to use.
//
func (self *Squeezer) SetSwitchThresholdPolicy(policy FoldThresholdPolicy) {
	var _arg0 *C.AdwSqueezer           // out
	var _arg1 C.AdwFoldThresholdPolicy // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwFoldThresholdPolicy(policy)

	C.adw_squeezer_set_switch_threshold_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetTransitionDuration sets the transition animation duration for self.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (self *Squeezer) SetTransitionDuration(duration uint) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.guint        // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.adw_squeezer_set_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation used for transitions between
// children in self.
//
// The function takes the following parameters:
//
//    - transition: new transition type.
//
func (self *Squeezer) SetTransitionType(transition SqueezerTransitionType) {
	var _arg0 *C.AdwSqueezer              // out
	var _arg1 C.AdwSqueezerTransitionType // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwSqueezerTransitionType(transition)

	C.adw_squeezer_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(transition)
}

// SetXAlign sets the horizontal alignment, from 0 (start) to 1 (end).
//
// The function takes the following parameters:
//
//    - xalign: new alignment value.
//
func (self *Squeezer) SetXAlign(xalign float32) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.float        // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(xalign)

	C.adw_squeezer_set_xalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the vertical alignment, from 0 (top) to 1 (bottom).
//
// The function takes the following parameters:
//
//    - yalign: new alignment value.
//
func (self *Squeezer) SetYAlign(yalign float32) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.float        // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(yalign)

	C.adw_squeezer_set_yalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(yalign)
}

// SqueezerPage: auxiliary class used by adw.Squeezer.
type SqueezerPage struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*SqueezerPage)(nil)
)

func wrapSqueezerPage(obj *externglib.Object) *SqueezerPage {
	return &SqueezerPage{
		Object: obj,
	}
}

func marshalSqueezerPager(p uintptr) (interface{}, error) {
	return wrapSqueezerPage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child returns the squeezer child to which self belongs.
func (self *SqueezerPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwSqueezerPage // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.AdwSqueezerPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_page_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Enabled gets whether self is enabled.
func (self *SqueezerPage) Enabled() bool {
	var _arg0 *C.AdwSqueezerPage // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSqueezerPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_squeezer_page_get_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEnabled sets whether self is enabled.
//
// The function takes the following parameters:
//
//    - enabled: whether self is enabled.
//
func (self *SqueezerPage) SetEnabled(enabled bool) {
	var _arg0 *C.AdwSqueezerPage // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSqueezerPage)(unsafe.Pointer(self.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_page_set_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enabled)
}
