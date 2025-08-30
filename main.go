package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
)

func main() {
	app := adw.NewApplication("com.github.diamondburned.gotk4-examples.gtk4.simple", gio.ApplicationFlagsNone)

	app.ConnectActivate(func() { activate(app) })
	
	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func homefab() *gtk.Button {
	fab := gtk.NewButton()
	fab.SetLabel("  Create new shortcut")
	fab.SetMarginBottom(10)
	fab.SetMarginEnd(10)
	fab.SetHAlign(gtk.AlignEnd)
	fab.SetVAlign(gtk.AlignEnd)
	fab.AddCSSClass("suggested-action")
	fab.AddCSSClass("circular")
	fab.AddCSSClass("fab")
	return fab
}

func homePage(app *adw.Application) *gtk.Overlay {
	overlay := gtk.NewOverlay()
	overlay.SetHExpand(true)
	overlay.SetVExpand(true)

	grid := gtk.NewGrid()
	overlay.SetChild(grid)

	fab := homefab()
	

	overlay.AddOverlay(fab)
	
	return overlay
}

func activate(app *adw.Application) {
	cssProvider := gtk.NewCSSProvider()
	cssProvider.LoadFromPath("./static/style.css")

	gtk.StyleContextAddProviderForDisplay(
		gdk.DisplayGetDefault(),
		cssProvider,
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

	window := adw.NewApplicationWindow(&app.Application)
	window.SetTitle("GameCutter")
	window.SetDefaultSize(600, 400)

	stack := gtk.NewStack()
	stack.SetTransitionType(gtk.StackTransitionTypeSlideLeftRight)

	
	stack.AddTitled( homePage(app) , "homepage", "Home")
	
	window.SetContent(stack)

	window.Show()
}