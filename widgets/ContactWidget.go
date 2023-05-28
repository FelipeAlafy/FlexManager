package widgets

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func PreFormContactWhatsapp(name string, box *gtk.Box) (*gtk.Entry, *gtk.Button) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	entry, err := gtk.EntryNew()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)
	entry.SetPlaceholderText("DDD+Numero")

	MustUseText := false
	pixbuf, err := gdk.PixbufNewFromFileAtScale("resources/whatsapp.svg", 30, -1, true)
	if err != nil {MustUseText = true}
	image, err := gtk.ImageNewFromPixbuf(pixbuf)
	if err != nil {MustUseText = true}

	whatsapp, err := gtk.ButtonNew()
	handler.Error("widgets/ContactWidget.go >> PreFormContactWhatsapp >> whatsapp before if", err)
	if MustUseText {
		whatsapp, err = gtk.ButtonNewWithLabel("Whatsapp")
		handler.Error("widgets/ContactWidget.go >> PreFormContactWhatsapp >> whatsapp on if MustUseText was true", err)
	} 
	whatsapp.SetImage(image)

	form.PackStart(lbl, true, true, 10)
	form.PackStart(entry, true, true, 20)
	form.PackStart(whatsapp, false, false, 5)

	box.PackStart(form, true, true, 10)

	return entry, whatsapp
}