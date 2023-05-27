package widgets

import (
	"strings"
	"unicode"

	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/glib"
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

func PreFormForPay(form *gtk.Box, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION int) (*gtk.ListStore, *gtk.Entry, *gtk.ComboBoxText, *gtk.Entry, *gtk.Button) {
	frame, err := gtk.FrameNew("Pagamento")
	handler.Error("ui/widgets.go >> PreFormForPay >> frame", err)
	
	valor, err := gtk.EntryNew()
	handler.Error("ui/widgets.go >> PreFormForPay >> valor", err)
	valor.SetPlaceholderText("Digite o valor")

	obs, _ := gtk.EntryNew()
	handler.Error("ui/widgets.go >> PreFormForPay >> obs", err)
	obs.SetPlaceholderText("Observação do pagamento")
	obs.SetTooltipText("Este campo existe para que\nsejam colocadas observações\nExemplo: Coloque em quantas parcelas o\nprojeto foi feita.")
	
	valor.Connect("changed", func ()  {
		s, _ := valor.GetText()
		for _, c := range s {
			if c == ',' {continue}
			if c == '.' {continue}
			if unicode.IsDigit(c) {continue}
			valor.SetText(strings.Replace(s, string(c), "", 1))
		}
	})

	payBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	headerPayBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)

	payCombo := PreFormComboBox("Forma de Pagemento: ", []string{"Dinheiro", "Cartão Crédito", "Cartão Débito", "Pix", "Boleto", "Cheque", "Depósito"}, headerPayBox)
	add, _ := gtk.ButtonNewFromIconName("list-add-symbolic", gtk.ICON_SIZE_BUTTON)
	
	headerPayBox.PackStart(valor, false, false, 10)
	headerPayBox.PackStart(payCombo, false, false, 10)
	headerPayBox.PackStart(obs, true, true, 10)
	headerPayBox.PackStart(add, false, false, 10)
	payBox.PackStart(headerPayBox, true, true, 5)

	//Tree
	storage, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	handler.Error("ui/widgets.go >> PreFormForPay >> storage", err)
	
	tree, remove := setupTreeView(storage, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION)
	headerPayBox.PackStart(remove, false, false, 0)
	payBox.PackStart(tree, false, true, 10)

	frame.Add(payBox)
	form.PackStart(frame, true, true, 10)
	return storage, valor, payCombo, obs, add
}

func createColumn(name string, id int) *gtk.TreeViewColumn {
	render, err := gtk.CellRendererTextNew()
	handler.Error("ui/Widgets.go >> createColumn >> render", err)

	column, err := gtk.TreeViewColumnNewWithAttribute(name, render, "text", id)
	handler.Error("ui/Widgets.go >> createColumn >> column", err)
	return column
}

func setupTreeView(listStore *gtk.ListStore, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION int) (*gtk.TreeView, *gtk.Button) {
	tree, err := gtk.TreeViewNew()
	handler.Error("ui/Widgets.go >> setupTreeView >> tree", err)

	tree.AppendColumn(createColumn("Tipo de Pagamento", 0))
	tree.AppendColumn(createColumn("Valor", 1))
	tree.AppendColumn(createColumn("Observações", 2))

	tree.SetModel(listStore)
	tree.SetHExpand(true)

	// getValue := func (c *glib.Value) string {
	// 	s, err := c.GetString()
	// 	if err != nil {return ""}
	// 	return s
	// }

	remove, _ := gtk.ButtonNewFromIconName("app-remove-symbolic", gtk.ICON_SIZE_BUTTON)

	selected, _ := tree.GetSelection()
	selected.SetMode(gtk.SELECTION_SINGLE)
	remove.Connect("clicked", func () {
		_, selectedIter, ok := selected.GetSelected()
		if !ok {return}
		
		//Data From Selected Iter
		if !listStore.IterIsValid(selectedIter) {return}
		payType1, _ := listStore.GetValue(selectedIter, 0)
		value1, _ := listStore.GetValue(selectedIter, 1)
		pt1S, _ := payType1.GetString()
		v1S, err := value1.GetString()
		if err != nil {return}

		listStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			//Data From Next Iter
			payType2, _ := listStore.GetValue(iter, 0)
			value2, _ := listStore.GetValue(iter, 1)
			pt2S, _ := payType2.GetString()
			v2S, _ := value2.GetString()

			if pt1S == pt2S && v1S == v2S {
				listStore.Remove(iter)
				selected.UnselectAll()
				pt1S = ""
				v1S = ""
				return true
			}
			return false
		})
	})
	

	return tree, remove
}

func AddRow(list *gtk.ListStore, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION int, payType, value, obs string) {
	iter := list.Append()

	err := list.Set(iter, []int{0, 1, 2},
					[]interface{}{payType, value, obs})
	handler.Error("ui/widgets.go >> AddRow >> err", err)
}