package widgets

import (
	"strings"
	"unicode"

	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	COLUMN_EXPENSE_TYPE = iota
	COLUMN_EXPENSE_VALUE
	COLUMN_EXPENSE_OBSERVATION
)

func PreFormLabel(name string, form *gtk.Box) *gtk.Label {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("widgets/Financial.go >> PreFormLabel >> box", err)

	label, err := gtk.LabelNew(name + ":")
	handler.Error("widgets/Financial.go >> PreFormLabel >> label", err)

	valLabel, err := gtk.LabelNew("")
	handler.Error("widgets/Financial.go >> PreFormLabel >> box", err)
	box.PackStart(label, true, true, 0)
	box.PackStart(valLabel, true, true, 0)
	form.PackStart(box, false, false, 5)
	return valLabel
}

func PreFormTreeFinancial(form *gtk.Box, vl *gtk.Label) (*gtk.ListStore, *gtk.ComboBoxText, *gtk.Entry, *gtk.Entry, *gtk.Button, *gtk.Button) {
	frame, err := gtk.FrameNew("Cadastro de Despesas")
	handler.Error("widgets/Financial.go >> PreFormTree >> frame", err)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("widgets/Financial.go >> PreFormTree >> box", err)
	
	topBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("widgets/Financial.go >> PreFormTree >> topBox", err)

	lblType, err := gtk.LabelNew("Tipo de despesa: ")
	handler.Error("widgets/Financial.go >> PreFormTree >> lblType", err)
	lblValue, err := gtk.LabelNew("Valor: ")
	handler.Error("widgets/Financial.go >> PreFormTree >> lblValue", err)
	lblObs, err := gtk.LabelNew("Observação: ")
	handler.Error("widgets/Financial.go >> PreFormTree >> lblObs", err)

	expenseType, err := gtk.ComboBoxTextNew()
	handler.Error("widgets/Financia.go >> PreFormTree >> expensesType", err)
	
	options := []string{"Alimentação", "Material de escritório", "Transporte", "Despesa Geral", "Despesa Recorrente", "Outro"}

	for _, o := range options {
		expenseType.AppendText(o)
	}

	expenseValue, err := gtk.EntryNew()
	handler.Error("widgets/Financia.go >> PreFormTree >> expensesType", err)

	expenseObservation, err := gtk.EntryNew()
	handler.Error("widgets/Financia.go >> PreFormTree >> expensesObservation", err)


	add, _ := gtk.ButtonNewFromIconName("list-add-symbolic", gtk.ICON_SIZE_BUTTON)

	//Tree 
	storage, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	handler.Error("widgets/Financial.go >> PreFormTree >> storage", err)

	tree, remove := setupTreeViewExpense(storage, COLUMN_EXPENSE_TYPE, COLUMN_EXPENSE_VALUE, COLUMN_EXPENSE_OBSERVATION, vl)

	expenseValue.Connect("changed", func ()  {
		s, _ := expenseValue.GetText()
		for _, c := range s {
			if c == ',' {continue}
			if c == '.' {continue}
			if unicode.IsDigit(c) {continue}
			expenseValue.SetText(strings.Replace(s, string(c), "", 1))
		}
	})

	topBox.PackStart(lblType, false, false, 5)
	topBox.PackStart(expenseType, false, false, 10)
	topBox.PackStart(lblValue, false, false, 5)
	topBox.PackStart(expenseValue, false, false, 10)
	topBox.PackStart(lblObs, false, false, 5)
	topBox.PackStart(expenseObservation, false, false, 10)
	topBox.PackStart(add, false, false, 10)
	topBox.PackStart(remove, false, false, 10)

	box.PackStart(topBox, false, false, 10)
	box.PackStart(tree, true, true, 10)
	frame.Add(box)
	form.PackStart(frame, true, true, 0)
	
	return storage, expenseType, expenseValue, expenseObservation, add, remove
}

func setupTreeViewExpense(listStore *gtk.ListStore, COLUMN_EXPENSE_TYPE, COLUMN_EXPENSE_VALUE, COLUMN_EXPENSE_OBSERVATION int, vl *gtk.Label) (*gtk.TreeView, *gtk.Button) {
	l, err := vl.GetText()
	if err != nil {l = "0,00"}
	var totalValue = handler.ConvertStringIntoFloat(l)

	tree, err := gtk.TreeViewNew()
	handler.Error("ui/Widgets.go >> setupTreeView >> tree", err)

	tree.AppendColumn(createColumn("Tipo de Pagamento", COLUMN_EXPENSE_TYPE))
	tree.AppendColumn(createColumn("Valor", COLUMN_EXPENSE_VALUE))
	column := createColumn("Observação", COLUMN_EXPENSE_OBSERVATION)
	column.SetExpand(true)
	tree.AppendColumn(column)

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
				vl.SetText(handler.GetCashFormatted(totalValue))

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

func AddRowFinancial(list *gtk.ListStore, expenseType, expenseValue, expenseObservation string, vl *gtk.Label) {
	iter := list.Append()

	err := list.Set(iter, []int{0, 1, 2},
					[]interface{}{expenseType, expenseValue, expenseObservation})
	handler.Error("ui/widgets.go >> AddRow >> err", err)
	str, err := vl.GetText()
	if err != nil {str = "0,00"}
	totalValue := handler.ConvertStringIntoFloat(str)
	totalValue += handler.ConvertStringIntoFloat(expenseValue)
	vl.SetText(handler.GetCashFormatted(totalValue))
}