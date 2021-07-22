package main

import (
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-examples.gtk4.simple", 0)
	app.Connect("activate", activate)

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	stack := gtk.NewStack()
	stack.SetHExpand(true)
	stack.SetSizeRequest(200, -1)
	stack.SetTransitionType(gtk.StackTransitionTypeSlideUpDown)
	stack.AddTitled(labelPage(), "label", "Label")
	stack.AddTitled(buttonPage(), "button", "Button")

	switcher := gtk.NewStackSidebar()
	switcher.SetSizeRequest(200, -1)
	switcher.SetStack(stack)

	flap := adw.NewFlap()
	flap.SetFlap(switcher)
	flap.SetContent(stack)
	flap.SetSwipeToOpen(true)
	flap.SetSwipeToClose(true)
	flap.SetFoldPolicy(adw.FlapFoldPolicyAuto)
	flap.SetTransitionType(adw.FlapTransitionTypeOver)
	flap.SetSeparator(gtk.NewSeparator(gtk.OrientationVertical))

	stack.InitiallyUnowned.Connect("notify::visible-child", func() {
		// Collapse the flap if we're in mobile view and the flap is opened.
		if flap.Folded() && flap.RevealFlap() {
			flap.SetRevealFlap(false)
		}
	})

	unflap := gtk.NewButtonFromIconName("document-properties-symbolic")
	unflap.InitiallyUnowned.Connect("clicked", func() {
		flap.SetRevealFlap(!flap.RevealFlap())
	})

	header := gtk.NewHeaderBar()
	header.SetShowTitleButtons(true)
	header.PackStart(unflap)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Adwaita Example")
	window.SetChild(flap)
	window.SetTitlebar(header)
	window.SetDefaultSize(600, 300)
	window.Show()
}

func labelPage() gtk.Widgetter {
	label := gtk.NewLabel("")
	label.SetMarkup(`<span size="xx-large">Hello, Go!</span>`)
	return label
}

func buttonPage() gtk.Widgetter {
	info := gtk.NewInfoBar()
	info.SetVAlign(gtk.AlignStart)
	info.SetHAlign(gtk.AlignCenter)
	info.AddChild(gtk.NewLabel("Button clicked."))
	info.SetShowCloseButton(true)
	info.SetMessageType(gtk.MessageInfo)
	info.Hide()
	info.InitiallyUnowned.Connect("response", (*gtk.InfoBar).Hide)

	button := gtk.NewButtonWithLabel("Click me!")
	button.SetVAlign(gtk.AlignCenter)
	button.SetHAlign(gtk.AlignCenter)
	button.InitiallyUnowned.Connect("clicked", func() { info.Show() })

	overlay := gtk.NewOverlay()
	overlay.SetHExpand(true)
	overlay.SetVExpand(true)
	overlay.SetChild(button)
	overlay.AddOverlay(info)

	return overlay
}
