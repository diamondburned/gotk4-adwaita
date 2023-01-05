package main

//go:generate go run . -o ./pkg/

import (
	"path"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
)

func main() {
	genmain.Run(Data)
}

const (
	gotk4Module   = "github.com/diamondburned/gotk4/pkg"
	adwaitaModule = "github.com/diamondburned/gotk4-adwaita/pkg"
)

var Data = genmain.Data{
	Module: adwaitaModule,
	ExternOverrides: map[string]gir.Repositories{
		gotk4Module: genmain.MustLoadPackages(gendata.Packages),
	},
	Packages: []genmain.Package{
		{Name: "libadwaita-1"},
	},
	PkgGenerated: []string{"adw"},
	PkgExceptions: []string{
		"go.mod",
		"go.sum",
		"LICENSE",
		"_examples",
	},
	Postprocessors: map[string][]girgen.Postprocessor{
		"Adw-1": {AdwInitPreserveTheme},
	},
}

func AdwInitPreserveTheme(nsgen *girgen.NamespaceGenerator) error {
	fg := nsgen.MakeFile("adw-init.go")

	h := fg.Header()
	h.Import(path.Join(gotk4Module, "gtk", "v4"))

	p := fg.Pen()
	p.Line(`
		// Init calls adw.Init then restores the user's theme preference. This
		// function prevents adw.Init from overriding the user's theme to
		// Adwaita.
		//
		// For more information, see
		// https://gitlab.gnome.org/GNOME/libadwaita/-/issues/215.
		func InitPreserveTheme() {
			Init()
			settings := gtk.SettingsGetDefault()
			settings.ResetProperty("gtk-theme-name")
		}
	`)

	return nil
}
