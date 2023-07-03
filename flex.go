package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"

	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/ui"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(">> Error while loading .env file")
	}

	db := database.Run(os.Getenv("FLEXMANAGERUSER"), os.Getenv("FLEXMANAGERPASS"), os.Getenv("FLEXMANAGERHOST"), os.Getenv("FLEXMANAGERPORT"), os.Getenv("FLEXMANAGERDB"))

	app, err := gtk.ApplicationNew("com.github.FelipeAlafy.FlexManager", glib.APPLICATION_FLAGS_NONE)
	handler.Error("ui/flex.go >> Line 10", err)
	app.Connect("activate", func() { ui.OnActivate(app, db) })
	os.Exit(app.Run(os.Args))
}
