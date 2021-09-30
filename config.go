package main

import (
	"path"

	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module   = "github.com/diamondburned/gotk4/pkg"
	adwaitaModule = "github.com/diamondburned/gotk4-adwaita/pkg"
)

var packages = []gendata.Package{
	{PkgName: "libadwaita-1", Namespaces: nil},
}

var pkgGenerated = []string{
	"adw",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
	"_examples",
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

var postprocessors = map[string][]girgen.Postprocessor{
	"Adw-1": {AdwInitPreserveTheme},
}

var filters = []types.FilterMatcher{}
