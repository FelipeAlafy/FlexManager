package main

import (
	"os"

	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/ui"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	//Runnable folder
	os.Mkdir("Runnable", 0700)

	app, err := gtk.ApplicationNew("com.github.FelipeAlafy.FlexManager", glib.APPLICATION_FLAGS_NONE)
	handler.Error("ui/flex.go >> Line 10", err)
	app.Connect("activate", func() { ui.OnActivate(app) })
	os.Exit(app.Run(os.Args))
}
