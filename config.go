package main

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
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

var filters = []types.FilterMatcher{}
