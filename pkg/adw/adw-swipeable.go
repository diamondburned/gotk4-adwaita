// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.adw_swipeable_get_type()), F: marshalSwipeabler},
	})
}

// SwipeableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SwipeableOverrider interface {
	// CancelProgress gets the progress self will snap back to after the gesture
	// is canceled.
	CancelProgress() float64
	// Distance gets the swipe distance of self.
	//
	// This corresponds to how many pixels 1 unit represents.
	Distance() float64
	// Progress gets the current progress of self.
	Progress() float64
	// SnapPoints gets the snap points of self.
	//
	// Each snap point represents a progress value that is considered acceptable
	// to end the swipe on.
	SnapPoints() []float64
	// SwipeArea gets the area self can start a swipe from for the given
	// direction and gesture type.
	//
	// This can be used to restrict swipes to only be possible from a certain
	// area, for example, to only allow edge swipes, or to have a draggable
	// element and ignore swipes elsewhere.
	//
	// If not implemented, the default implementation returns the allocation of
	// self, allowing swipes from anywhere.
	SwipeArea(navigationDirection NavigationDirection, isDrag bool) *gdk.Rectangle
}

// Swipeable: interface for swipeable widgets.
//
// The AdwSwipeable interface is implemented by all swipeable widgets.
//
// See adw.SwipeTracker for details about implementing it.
type Swipeable struct {
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*Swipeable)(nil)
)

// Swipeabler describes Swipeable's interface methods.
type Swipeabler interface {
	externglib.Objector

	// CancelProgress gets the progress self will snap back to after the gesture
	// is canceled.
	CancelProgress() float64
	// Distance gets the swipe distance of self.
	Distance() float64
	// Progress gets the current progress of self.
	Progress() float64
	// SnapPoints gets the snap points of self.
	SnapPoints() []float64
	// SwipeArea gets the area self can start a swipe from for the given
	// direction and gesture type.
	SwipeArea(navigationDirection NavigationDirection, isDrag bool) *gdk.Rectangle
}

var _ Swipeabler = (*Swipeable)(nil)

func wrapSwipeable(obj *externglib.Object) *Swipeable {
	return &Swipeable{
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

func marshalSwipeabler(p uintptr) (interface{}, error) {
	return wrapSwipeable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CancelProgress gets the progress self will snap back to after the gesture is
// canceled.
func (self *Swipeable) CancelProgress() float64 {
	var _arg0 *C.AdwSwipeable // out
	var _cret C.double        // in

	_arg0 = (*C.AdwSwipeable)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipeable_get_cancel_progress(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Distance gets the swipe distance of self.
//
// This corresponds to how many pixels 1 unit represents.
func (self *Swipeable) Distance() float64 {
	var _arg0 *C.AdwSwipeable // out
	var _cret C.double        // in

	_arg0 = (*C.AdwSwipeable)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipeable_get_distance(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Progress gets the current progress of self.
func (self *Swipeable) Progress() float64 {
	var _arg0 *C.AdwSwipeable // out
	var _cret C.double        // in

	_arg0 = (*C.AdwSwipeable)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipeable_get_progress(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SnapPoints gets the snap points of self.
//
// Each snap point represents a progress value that is considered acceptable to
// end the swipe on.
func (self *Swipeable) SnapPoints() []float64 {
	var _arg0 *C.AdwSwipeable // out
	var _cret *C.double       // in
	var _arg1 C.int           // in

	_arg0 = (*C.AdwSwipeable)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipeable_get_snap_points(_arg0, &_arg1)
	runtime.KeepAlive(self)

	var _gdoubles []float64 // out

	defer C.free(unsafe.Pointer(_cret))
	_gdoubles = make([]float64, _arg1)
	copy(_gdoubles, unsafe.Slice((*float64)(unsafe.Pointer(_cret)), _arg1))

	return _gdoubles
}

// SwipeArea gets the area self can start a swipe from for the given direction
// and gesture type.
//
// This can be used to restrict swipes to only be possible from a certain area,
// for example, to only allow edge swipes, or to have a draggable element and
// ignore swipes elsewhere.
//
// If not implemented, the default implementation returns the allocation of
// self, allowing swipes from anywhere.
//
// The function takes the following parameters:
//
//    - navigationDirection: direction of the swipe.
//    - isDrag: whether the swipe is caused by a dragging gesture.
//
func (self *Swipeable) SwipeArea(navigationDirection NavigationDirection, isDrag bool) *gdk.Rectangle {
	var _arg0 *C.AdwSwipeable          // out
	var _arg1 C.AdwNavigationDirection // out
	var _arg2 C.gboolean               // out
	var _arg3 C.GdkRectangle           // in

	_arg0 = (*C.AdwSwipeable)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwNavigationDirection(navigationDirection)
	if isDrag {
		_arg2 = C.TRUE
	}

	C.adw_swipeable_get_swipe_area(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(navigationDirection)
	runtime.KeepAlive(isDrag)

	var _rect *gdk.Rectangle // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _rect
}
