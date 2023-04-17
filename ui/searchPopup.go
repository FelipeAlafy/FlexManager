package ui

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)


func popup(root *gtk.SearchEntry) (*gtk.Box) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/SearchPopup.go >> box", err)
	
	pop, err := gtk.PopoverNew(root)
	handler.Error("ui/SearchPopup.go >> pop", err)
	pop.SetPosition(gtk.POS_BOTTOM)
	pop.Add(box)
	pop.ShowAll()
	return box
}