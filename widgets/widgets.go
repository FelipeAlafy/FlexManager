package widgets

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

func PreForm(name string, box *gtk.Box) *gtk.Entry {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	entry, err := gtk.EntryNew()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)

	form.PackStart(lbl, true, true, 10)
	form.PackEnd(entry, true, true, 20)

	box.PackStart(form, true, true, 10)

	return entry
}

func PreFormComboBox(name string, options []string, box *gtk.Box) (*gtk.ComboBoxText) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	entry, err := gtk.ComboBoxTextNew()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)

	for _, element := range options {
		entry.AppendText(element)
	} 

	form.PackStart(lbl, true, false, 10)
	form.PackEnd(entry, true, true, 20)

	box.PackStart(form, true, false, 10)

	return entry
}

func PreFormComboBoxWithAnEntry(name string, options []string, box *gtk.Box) (*gtk.ComboBoxText) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	entry, err := gtk.ComboBoxTextNewWithEntry()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)

	for _, element := range options {
		entry.AppendText(element)
	} 

	form.PackStart(lbl, true, false, 10)
	form.PackEnd(entry, true, true, 20)

	box.PackStart(form, true, false, 10)

	return entry
}

func PreFormCalendar(name string, box *gtk.Box) (*gtk.Calendar) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	entry, err := gtk.CalendarNew()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)

	form.PackStart(lbl, true, false, 10)
	form.PackEnd(entry, true, true, 20)

	box.PackStart(form, true, false, 10)

	return entry
}

func PreFormTextView(name string, box *gtk.Box) (*gtk.TextView) {
	frame, _ := gtk.FrameNew("")
	scrollable, _ := gtk.ScrolledWindowNew(nil, nil)

	scrollable.SetHExpand(true)
	scrollable.SetBorderWidth(3)

	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preFormTextView >> gtk.Box, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preFormTextView >> gtk.Label, while creating a label for "+name, err)

	view, err := gtk.TextViewNew()
	handler.Error("ui/widgets.go >> preFormTextView >> gtk.TextView, while creating an Entry for "+name, err)
	
	scrollable.Add(view)
	frame.Add(scrollable)

	form.PackStart(lbl, true, false, 10)
	form.PackEnd(frame, true, true, 20)

	box.PackStart(form, true, false, 10)

	return view
}

func PreFormCheckBox(name string, box *gtk.Box) *gtk.CheckButton {
	form, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/widgets.go >> preForm >> Line 46, while creating a form for "+name, err)

	lbl, err := gtk.LabelNew(name + ": ")
	handler.Error("ui/widgets.go >> preForm >> Line 51, while creating a label for "+name, err)

	entry, err := gtk.CheckButtonNew()
	handler.Error("ui/widgets.go >> preForm >> Line 53, while creating an Entry for "+name, err)

	form.PackStart(lbl, true, true, 10)
	form.PackEnd(entry, true, true, 20)

	box.PackStart(form, true, true, 10)

	return entry
}