// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSpringParams = coreglib.Type(C.adw_spring_params_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpringParams, F: marshalSpringParams},
	})
}

// SpringParams: physical parameters of a spring for springanimation.
//
// Any spring can be described by three parameters: mass, stiffness and damping.
//
// An undamped spring will produce an oscillatory motion which will go on
// forever.
//
// The frequency and amplitude of the oscillations will be determined by the
// stiffness (how "strong" the spring is) and its mass (how much "inertia" it
// has).
//
// If damping is larger than 0, the amplitude of that oscillating motion will
// exponientally decrease over time. If that damping is strong enough that the
// spring can't complete a full oscillation, it's called an overdamped spring.
//
// If we the spring can oscillate, it's called an underdamped spring.
//
// The value between these two behaviors is called critical damping;
// a critically damped spring will comes to rest in the minimum possible time
// without producing oscillations.
//
// The damping can be replaced by damping ratio, which produces the following
// springs:
//
// * 0: an undamped spring. * Between 0 and 1: an underdamped spring. * 1:
// a critically damped spring. * Larger than 1: an overdamped spring.
//
// # As such
//
// An instance of this type is always passed by reference.
type SpringParams struct {
	*springParams
}

// springParams is the struct that's finalized.
type springParams struct {
	native *C.AdwSpringParams
}

func marshalSpringParams(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SpringParams{&springParams{(*C.AdwSpringParams)(b)}}, nil
}

// NewSpringParams constructs a struct SpringParams.
func NewSpringParams(dampingRatio float64, mass float64, stiffness float64) *SpringParams {
	var _arg1 C.double           // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _cret *C.AdwSpringParams // in

	_arg1 = C.double(dampingRatio)
	_arg2 = C.double(mass)
	_arg3 = C.double(stiffness)

	_cret = C.adw_spring_params_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(dampingRatio)
	runtime.KeepAlive(mass)
	runtime.KeepAlive(stiffness)

	var _springParams *SpringParams // out

	_springParams = (*SpringParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_springParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_spring_params_unref((*C.AdwSpringParams)(intern.C))
		},
	)

	return _springParams
}

// NewSpringParamsFull constructs a struct SpringParams.
func NewSpringParamsFull(damping float64, mass float64, stiffness float64) *SpringParams {
	var _arg1 C.double           // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _cret *C.AdwSpringParams // in

	_arg1 = C.double(damping)
	_arg2 = C.double(mass)
	_arg3 = C.double(stiffness)

	_cret = C.adw_spring_params_new_full(_arg1, _arg2, _arg3)
	runtime.KeepAlive(damping)
	runtime.KeepAlive(mass)
	runtime.KeepAlive(stiffness)

	var _springParams *SpringParams // out

	_springParams = (*SpringParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_springParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_spring_params_unref((*C.AdwSpringParams)(intern.C))
		},
	)

	return _springParams
}

// Damping gets the damping of self.
//
// The function returns the following values:
//
//   - gdouble: damping.
//
func (self *SpringParams) Damping() float64 {
	var _arg0 *C.AdwSpringParams // out
	var _cret C.double           // in

	_arg0 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = C.adw_spring_params_get_damping(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DampingRatio gets the damping ratio of self.
//
// The function returns the following values:
//
//   - gdouble: damping ratio.
//
func (self *SpringParams) DampingRatio() float64 {
	var _arg0 *C.AdwSpringParams // out
	var _cret C.double           // in

	_arg0 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = C.adw_spring_params_get_damping_ratio(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Mass gets the mass of self.
//
// The function returns the following values:
//
//   - gdouble: mass.
//
func (self *SpringParams) Mass() float64 {
	var _arg0 *C.AdwSpringParams // out
	var _cret C.double           // in

	_arg0 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = C.adw_spring_params_get_mass(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Stiffness gets the stiffness of self.
//
// The function returns the following values:
//
//   - gdouble: stiffness.
//
func (self *SpringParams) Stiffness() float64 {
	var _arg0 *C.AdwSpringParams // out
	var _cret C.double           // in

	_arg0 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = C.adw_spring_params_get_stiffness(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}