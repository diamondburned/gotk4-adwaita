package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const adwaitaModule = "github.com/diamondburned/gotk4-adwaita/pkg"

var preprocessors = []types.Preprocessor{
	// We'll add our own AddSetter that is a bit more Go-like.
	types.RenameCallable("Adw-1.Breakpoint.add_setter", "add_setter_direct"),
}

var postprocessors = map[string][]girgen.Postprocessor{
	"Adw-1": {betterBreakpointAddSetter},
}

var Data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: adwaitaModule,
		Packages: []genmain.Package{
			{
				Name:       "libadwaita-1",
				Namespaces: []string{"Adw-1"},
			},
		},
		PkgGenerated: []string{"adw"},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
			"_examples",
		},
		Preprocessors:  preprocessors,
		Postprocessors: postprocessors,
	},
)

func betterBreakpointAddSetter(nsgen *girgen.NamespaceGenerator) error {
	fg := nsgen.MakeFile("adw-breakpoint-extras.go")
	fg.Header().Import(gendata.Module + "/glib/v2")

	p := fg.Pen()
	p.Line(`
		// AddSetter adds a setter to self.
		//
		// The setter will automatically set property on object to value when applying
		// the breakpoint, and set it back to its original value upon unapplying it.
		//
		// Note that setting properties to their original values does not work for
		// properties that have irreversible side effects. For example, changing
		// gtk.Button:label while gtk.Button:icon-name is set will reset the icon.
		// However, resetting the label will not set icon-name to its original value.
		//
		// Use the breakpoint::apply and breakpoint::unapply signals for those
		// properties instead.
		//
		// The function takes the following parameters:
		//
		//   - object: target object.
		//   - property: target property.
		//   - value to set.
		//
		func (self *Breakpoint) AddSetter(object glib.Objector, property string, value any) {
			self.AddSetterDirect(glib.BaseObject(object), property, glib.NewValue(value))
		}
	`)

	return nil
}

func main() {
	genmain.Run(Data)
}
