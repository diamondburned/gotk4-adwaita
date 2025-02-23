package main

import (
	adwaita "github.com/diamondburned/gotk4-adwaita"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

func main() {
	genmain.Run(adwaita.Data)
}
