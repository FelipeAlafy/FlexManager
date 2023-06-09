package widgets

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func PasswordEntry() *gtk.Entry {
	pass, err := gtk.EntryNew()
	handler.Error("ui/widgets.go >> PasswordEntry >> Pass", err)

	pass.SetVisibility(false)
	pass.SetIconFromIconName(gtk.ENTRY_ICON_SECONDARY, "view-reveal-symbolic.symbolic")
	pass.SetIconActivatable(gtk.ENTRY_ICON_SECONDARY, true)
	pass.Connect("icon-press", func ()  {
		isVisible := pass.GetVisibility()
		if isVisible {
			pass.SetVisibility(false)
			pass.SetIconFromIconName(gtk.ENTRY_ICON_SECONDARY, "view-reveal-symbolic.symbolic")
		} else {
			pass.SetVisibility(true)
			pass.SetIconFromIconName(gtk.ENTRY_ICON_SECONDARY, "view-conceal-symbolic.symbolic")
		}
	})
	return pass
}

func PreFormPassword(name string, box *gtk.Box) *gtk.Entry {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	pass := PasswordEntry()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)
	
	form.PackStart(lbl, true, true, 10)
	form.PackEnd(pass, true, true, 20)

	box.PackStart(form, true, true, 10)

	return pass
}

func PreFormPasswordEntryWithConfirmation(name string, box *gtk.Box) (*gtk.Entry, *gtk.Image) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("widgets/Password.go >> PreFormPasswordEntryWithConfirmation >> form", err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("widgets/Password.go >> PreFormPasswordEntryWithConfirmation >> lbl", err)

	pass := PasswordEntry()
	handler.Error("widgets/Password.go >> PreFormPasswordEntryWithConfirmation >> pass", err)

	pixbuf, err := gdk.PixbufNewFromFileAtScale("resources/checkround.svg", 32, 32, true)
	handler.Error("widgets/Password.go >> PreFormPasswordEntryWithConfirmation >> pixbuf", err)
	checker, err := gtk.ImageNewFromPixbuf(pixbuf)
	handler.Error("widgets/Password.go >> PreFormPasswordEntryWithConfirmation >> checker", err)
	checker.SetVisible(false)
	
	form.PackStart(lbl, true, true, 10)
	form.PackStart(pass, true, true, 20)
	form.PackStart(checker, false, false, 5)

	box.PackStart(form, true, true, 10)

	return pass, checker
}