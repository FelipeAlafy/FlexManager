package ui

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

func getHeaderbar() (*gtk.HeaderBar, []*gtk.Button, *gtk.SearchEntry) {
	bar, err := gtk.HeaderBarNew()
	handler.Error("ui/handler/Headerbar.go >> Line 10", err)

	bar.SetTitle("Flex Manager") //this should be dynamic by showing the name of the enterprise
	bar.SetShowCloseButton(true)

	client, err := gtk.ButtonNewFromIconName("contact-new-symbolic", gtk.ICON_SIZE_BUTTON)
	handler.Error("ui/handler/Headerbar.go >> Line 15", err)
	bar.PackStart(client)

	project, err := gtk.ButtonNewFromIconName("document-new-symbolic", gtk.ICON_SIZE_BUTTON)
	handler.Error("ui/handler/Headerbar.go >> Line 20", err)
	bar.PackStart(project)

	edit, err := gtk.ButtonNewFromIconName("document-edit-symbolic", gtk.ICON_SIZE_BUTTON)
	handler.Error("ui/handler/Headerbar.go >> Line 24", err)
	bar.PackStart(edit)

	searchbar, _ := gtk.SearchEntryNew()
	bar.PackEnd(searchbar)

	return bar, []*gtk.Button{client, project, edit}, searchbar
}
