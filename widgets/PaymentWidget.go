package widgets

import (
	"strings"
	"unicode"

	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func PreFormForPay(form *gtk.Box, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION int) (*gtk.ListStore, *gtk.Entry, *gtk.ComboBoxText, *gtk.Entry, *gtk.Button, *gtk.Button, *gtk.Label) {
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

	payCombo := PreFormComboBox("Forma de Pagemento", []string{"Dinheiro", "Cartão Crédito", "Cartão Débito", "Pix", "Boleto", "Cheque", "Depósito"}, headerPayBox)
	add, _ := gtk.ButtonNewFromIconName("list-add-symbolic", gtk.ICON_SIZE_BUTTON)
	
	headerPayBox.PackStart(valor, false, false, 10)
	headerPayBox.PackStart(payCombo, false, false, 10)
	headerPayBox.PackStart(obs, true, true, 10)
	headerPayBox.PackStart(add, false, false, 10)
	payBox.PackStart(headerPayBox, true, true, 5)

	//Bottom
	bottomBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("widgets/wigdets.go >> PreFormForPay >> bottomBox", err)
	l, err := gtk.LabelNew("Valor Total R$")
	handler.Error("widgets/wigdets.go >> PreFormForPay >> l", err)
	vl, err := gtk.LabelNew("0,00")
	handler.Error("widgets/wigdets.go >> PreFormForPay >> vl", err)
	bottomBox.PackStart(l, false, false, 10)
	bottomBox.PackStart(vl, false, false, 0)

	//Tree
	storage, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	handler.Error("ui/widgets.go >> PreFormForPay >> storage", err)
	
	tree, remove := setupTreeView(storage, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION, *vl)
	headerPayBox.PackStart(remove, false, false, 0)
	payBox.PackStart(tree, false, true, 10)

	payBox.PackStart(bottomBox, false, false, 5)
	frame.Add(payBox)
	form.PackStart(frame, true, true, 10)
	return storage, valor, payCombo, obs, add, remove, vl
}

func createColumn(name string, id int) *gtk.TreeViewColumn {
	render, err := gtk.CellRendererTextNew()
	handler.Error("ui/Widgets.go >> createColumn >> render", err)

	column, err := gtk.TreeViewColumnNewWithAttribute(name, render, "text", id)
	handler.Error("ui/Widgets.go >> createColumn >> column", err)
	return column
}

func setupTreeView(listStore *gtk.ListStore, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION int, vl gtk.Label) (*gtk.TreeView, *gtk.Button) {
	l, err := vl.GetText()
	if err != nil {l = "0,00"}
	var totalValue = handler.ConvertStringIntoFloat(l)
	tree, err := gtk.TreeViewNew()
	handler.Error("ui/Widgets.go >> setupTreeView >> tree", err)

	tree.AppendColumn(createColumn("Tipo de Pagamento", 0))
	tree.AppendColumn(createColumn("Valor", 1))
	tree.AppendColumn(createColumn("Observações", 2))

	tree.SetModel(listStore)
	tree.SetHExpand(true)

	pixbuf, _ := gdk.PixbufNewFromFileAtScale("resources/trash.svg", 20, 25, true)
	image, _ := gtk.ImageNewFromPixbuf(pixbuf)
	remove, _ := gtk.ButtonNew()
	remove.SetMarginEnd(10)
	remove.SetImage(image)

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
				strValor, _ := vl.GetText()
				totalValue = handler.ConvertStringIntoFloat(strValor)
				totalValue -= handler.ConvertStringIntoFloat(v1S)
				vl.SetText(handler.ConvertFloatIntoString(totalValue))

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

func AddRow(list *gtk.ListStore, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION int, payType, value, obs string, vl *gtk.Label) {
	iter := list.Append()

	err := list.Set(iter, []int{0, 1, 2},
					[]interface{}{payType, value, obs})
	handler.Error("ui/widgets.go >> AddRow >> err", err)
	str, err := vl.GetText()
	if err != nil {str = "0,00"}
	totalValue := handler.ConvertStringIntoFloat(str)
	totalValue += handler.ConvertStringIntoFloat(value)
	vl.SetText(handler.ConvertFloatIntoString(totalValue))
}