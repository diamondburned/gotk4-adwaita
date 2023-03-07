package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

const (
	gotk4Module   = "github.com/diamondburned/gotk4/pkg"
	adwaitaModule = "github.com/diamondburned/gotk4-adwaita/pkg"
)

var Data = genmain.Data{
	Module: adwaitaModule,
	ExternOverrides: map[string]gir.Repositories{
		gotk4Module: genmain.MustLoadPackages(gendata.Packages),
	},
	Packages:      []genmain.Package{{Name: "libadwaita-1"}},
	KnownPackages: gendata.Packages,
	PkgGenerated:  []string{"adw"},
	PkgExceptions: []string{
		"go.mod",
		"go.sum",
		"LICENSE",
		"_examples",
	},
}

func main() {
	genmain.Run(Data)
}
