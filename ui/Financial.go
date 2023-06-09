package ui

import (
	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

func Financial(dbFinancial *gorm.DB, notebook *gtk.Notebook, editButton *gtk.Button) *gtk.Box {
	mainBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 30)
	handler.Error("ui/Financial.go >> Financial >> mainBox", err)
	
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	handler.Error("ui/Financial.go >> Financial >> scrolledWindow", err)

	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Financial.go >> Financial >> form", err)

	topBar, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/Financial.go >> Financial >> topBar", err)
	
	backButton, err := gtk.ButtonNewFromIconName("go-previous-symbolic", gtk.ICON_SIZE_BUTTON)
	handler.Error("ui/Financial.go >> Financial >> backButton", err)
	date, err := gtk.LabelNew("Maio/2023")
	handler.Error("ui/Financial.go >> Financial >> date", err)
	forwardButton, err := gtk.ButtonNewFromIconName("go-next-symbolic", gtk.ICON_SIZE_BUTTON)
	handler.Error("ui/Financial.go >> Financial >> forwardButton", err)

	topBar.PackStart(backButton, false, false, 0)
	topBar.PackStart(date, true, true, 0)
	topBar.PackEnd(forwardButton, false, false, 0)

	projectsExpander, err := gtk.ExpanderNew("Detalhes de transações de projetos deste mês")
	handler.Error("ui/Financial.go >> Financial >> projectsExpander", err)
	projectsBox, allProjects, wipProjects, finishedProjects, grossIncoming1 := getProjectExpander()
	projectsExpander.Add(projectsBox)

	expensesExpander, err := gtk.ExpanderNew("Detalhes financeiros do mês")
	handler.Error("ui/Financial.go >> Financial >> expensesExpander", err)
	expensesBox, grossIncoming2, monthExpenses, netRevenue, storage := getExpensesExpander()
	expensesExpander.Add(expensesBox)

	controller.FinancialInit(dbFinancial, notebook, editButton, forwardButton, backButton, allProjects, wipProjects, finishedProjects, grossIncoming1, grossIncoming2, monthExpenses, netRevenue, storage, date)

	form.PackStart(projectsExpander, false, false, 10)
	form.PackStart(expensesExpander, false, false, 10)
	mainBox.PackStart(topBar, false, false, 0)
	scrolledWindow.Add(form)
	mainBox.PackStart(scrolledWindow, true, true, 0)
	return mainBox
}

func getProjectExpander() (*gtk.Box, *gtk.Label, *gtk.Label, *gtk.Label, *gtk.Label) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Financial.go >> projectExpander >> form", err)
	allProjects := widgets.PreFormLabel("Quantidade de projetos realizados", form)
	wipProjects := widgets.PreFormLabel("Quantidade de projetos em andamento", form)
	finishedProjects := widgets.PreFormLabel("Quantidade de projetos finalizados", form)
	grossIncoming := widgets.PreFormLabel("Receita bruta dos projetos", form)

	return form, allProjects, wipProjects, finishedProjects, grossIncoming
}

func getExpensesExpander() (*gtk.Box, *gtk.Label, *gtk.Label, *gtk.Label, *gtk.ListStore) {
	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Financial.go >> projectExpander >> form", err)
	grossIncoming := widgets.PreFormLabel("Receita bruta do mês", form)
	monthExpanses := widgets.PreFormLabel("Despesas do mês", form)
	netRevenue := widgets.PreFormLabel("Receita liquida do mês", form)
	grossIncoming.SetText("R$ 0,00")
	monthExpanses.SetText("R$ 0,00")
	netRevenue.SetText("R$ 0,00")

	//Tree
	storage, expenseType, expenseValue, expenseObservation, add, remove := widgets.PreFormTreeFinancial(form, monthExpanses)

	add.Connect("clicked", func ()  {
		expenseTypeStr := expenseType.GetActiveText()
		expenseValueStr, err := expenseValue.GetText()
		if err != nil {return}
		if expenseTypeStr == "" && expenseValueStr == "" {return}
		expenseObservationStr, err := expenseObservation.GetText()
		if err != nil {return}
		
		widgets.AddRowFinancial(storage, expenseTypeStr, expenseValueStr, expenseObservationStr, monthExpanses)

		gross, err  := grossIncoming.GetText()
		if err != nil && gross != "" {return}
		month, err := monthExpanses.GetText()
		if err != nil && gross != "" {return}

		value := (handler.ConvertStringIntoFloat(gross) - handler.ConvertStringIntoFloat(month))
		netRevenue.SetText(handler.GetCashFormatted(value))
	})

	remove.ConnectAfter("clicked", func ()  {
		gross, err  := grossIncoming.GetText()
		if err != nil && gross != "" {return}
		month, err := monthExpanses.GetText()
		if err != nil && gross != "" {return}
		value := (handler.ConvertStringIntoFloat(gross) - handler.ConvertStringIntoFloat(month))
		netRevenue.SetText(handler.GetCashFormatted(value))
	})

	return form, grossIncoming, monthExpanses, netRevenue, storage
}