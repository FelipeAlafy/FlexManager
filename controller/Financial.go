package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

type FinancialFields struct {
	db *gorm.DB
	notebook *gtk.Notebook
	edit *gtk.Button
	next *gtk.Button
	back *gtk.Button
	allProjects *gtk.Label
	wipProjects *gtk.Label
	finishedProjects *gtk.Label
	grossIncoming1 *gtk.Label
	grossIncoming2 *gtk.Label
	monthExpanses *gtk.Label
	netRevenue *gtk.Label
	storage		*gtk.ListStore
	date *gtk.Label
}

func FinancialInit(db *gorm.DB, notebook *gtk.Notebook, editButton, back, next *gtk.Button, AllProjects, WipProjects, FinishedProjects, GrossIncoming1, GrossIncoming2, MonthExpenses, NetRevenue *gtk.Label, storage *gtk.ListStore, time *gtk.Label) {
	fields := FinancialFields{db, notebook, editButton, back, next, AllProjects, WipProjects, FinishedProjects, GrossIncoming1, GrossIncoming2, MonthExpenses, NetRevenue, storage, time}
	thisPage := fields.notebook.GetNPages()

	notebook.Connect("page-removed", func (_ *gtk.Notebook, _ *gtk.Widget, pageRemoved uint)  {
		if pageRemoved < uint(thisPage) {thisPage = thisPage - 1}
	})

	notebook.Connect("switch-page", func (_ *gtk.Notebook, _ *gtk.Widget, index int)  {
		if thisPage != index {return}
		image, err := gtk.ImageNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
		handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
		editButton.SetImage(image)
	})

	fields.SetCurrentMonth()
	
	fields.loadData()

	fields.edit.Connect("clicked", func ()  {
		startDate, endDate := fields.GetDate()
		fields.Save(&startDate, &endDate)
		fields.Clear(false)
	})

	fields.back.Connect("clicked", func ()  {
		fields.Clear(true)
		fields.subtractMonth()
		fields.loadData()
	})

	fields.next.Connect("clicked", func ()  {
		fields.Clear(true)
		fields.addMonth()
		fields.loadData()
	})
}

func (fields FinancialFields) Save(startDate, endDate *time.Time) {
	model := getExpenseModel(fields)
	database.DeleteAllExpenses(fields.db, startDate, endDate)
	for _, e := range model {
		e.Save(fields.db)
	}
}

func (fields FinancialFields) loadData() {
	startDate, endDate := fields.GetDate()
	Payments := database.GetPaymentsDataByDate(fields.db, startDate, endDate)
	
	fields.grossIncoming1.SetText(handler.GetCashFormatted(getSum(Payments)))
	fields.grossIncoming2.SetText(handler.GetCashFormatted(getSum(Payments)))
	allProjects := database.GetCountOfProjects(fields.db, startDate, endDate)
	statusFinished := database.GetCountOfProjectsByStatus(fields.db, startDate, endDate, database.Project{Status: "Finalizado"})
	fields.allProjects.SetText(fmt.Sprint(allProjects))
	fields.wipProjects.SetText(fmt.Sprint(allProjects - statusFinished))
	fields.finishedProjects.SetText(fmt.Sprint(statusFinished))

	expenses := database.GetExpensesDataByDate(fields.db, startDate, endDate)

	sumExpenses := 0.0
	for _, e := range expenses {
		sumExpenses = sumExpenses + e.ExpenseValue
		widgets.AddRowFinancial(fields.storage, e.ExpenseType, handler.ConvertFloatIntoString(e.ExpenseValue), e.ExpenseObservation, fields.monthExpanses)
	}
	total, _ := fields.grossIncoming2.GetText()
	fields.netRevenue.SetText(handler.GetCashFormatted((handler.ConvertStringIntoFloat(total) - sumExpenses)))
}

func (fields FinancialFields) GetDate() (time.Time, time.Time) {
	dateStr, err := fields.date.GetText()
	handler.Error("controller/Financial.go >> GetDate >> dateStr", err)
	splitted := strings.Split(dateStr, "/")
	startMonth := time.Date(toInt(splitted[1]), getMonth(splitted[0]), 1, 0, 0, 0, 0, time.Local)
	endMonth := time.Date(toInt(splitted[1]), getMonth(splitted[0]), getLastDayOfMonth(startMonth), 0, 0, 0, 0, time.Local)
	return startMonth, endMonth 
}

func getMonth(month string) time.Month {
	switch (month) {
	case "Janeiro":
		return time.January
	case "Fevereiro":
		return time.February
	case "Março":
		return time.March
	case "Abril":
		return time.April
	case "Maio":
		return time.May
	case "Junho":
		return time.June
	case "Julho":
		return time.July
	case "Agosto":
		return time.August
	case "Setembro":
		return time.September
	case "Outubro":
		return time.October
	case "Novembro":
		return time.November
	case "Dezembro":
		return time.December
	}
	return time.January
}

func getMonthInPortuguese(month time.Month) string {
	switch (month) {
	case time.January:
		return "Janeiro"
	case time.February :
		return "Fevereiro"
	case time.March:
		return "Março"
	case time.April:
		return "Abril"
	case time.May:
		return "Maio"
	case time.June:
		return "Junho"
	case time.July:
		return "Julho"
	case time.August:
		return "Agosto"
	case time.September:
		return "Setembro"
	case time.October:
		return "Outubro"
	case time.November:
		return "Novembro"
	case time.December:
		return "Dezembro"
	}
	return "Janeiro"
}

func getLastDayOfMonth(date time.Time) int {
	switch date.Month() {
	case time.January:
		return 31
	case time.February:
		if date.Year() % 4 == 0 && date.Year() % 100 != 0 {
			return 29
		} else if date.Year() % 4 == 0 && date.Year() % 100 == 0 && date.Year() % 400 == 0 {
			return 29
		} else {
			return 28
		}
	case time.March:
		return 31
	case time.April:
		return 30
	case time.May:
		return 31
	case time.June:
		return 30
	case time.July:
		return 31
	case time.August:
		return 31
	case time.September:
		return 30
	case time.October:
		return 31
	case time.November:
		return 30
	case time.December:
		return 31
	}
	return 30
}

func getSum(array []database.Payment) float64 {
	var sum float64 = 0.0
	for _, p := range array {
		sum = sum + p.Value
	}
	return sum
}

func getExpenseModel(fields FinancialFields) []database.Expense {
	expenses := []database.Expense{}

	fields.storage.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		ExpenseType, err := model.GetValue(iter, 0)
		if err != nil {return false}
		ExpenseValue, err := model.GetValue(iter, 1)
		if err != nil {return false}
		ExpenseObservation, err := model.GetValue(iter, 2)
		if err != nil {return false}

		p, err := ExpenseType.GetString()
		if err != nil {return false}
		v, err := ExpenseValue.GetString()
		if err != nil {return false}
		o, err := ExpenseObservation.GetString()
		if err != nil {return false}
		
		if p != "" && v != "" {
			expense := database.Expense{ExpenseType: p, ExpenseValue: handler.ConvertStringIntoFloat(v), ExpenseObservation: o}
			date, _ := fields.GetDate()
			expense.CreatedAt = date 
			expenses = append(expenses, expense)
			return false
		}
		return true
	})
	return expenses
}

func (fields FinancialFields) Clear(skipMonth bool) {
	if !skipMonth {
		fields.SetCurrentMonth()
	}
	fields.allProjects.SetText("")
	fields.grossIncoming1.SetText("")
	fields.grossIncoming2.SetText("")
	fields.finishedProjects.SetText("")
	fields.monthExpanses.SetText("")
	fields.netRevenue.SetText("")
	fields.wipProjects.SetText("")
	fields.storage.Clear()
}

func (fields FinancialFields) SetCurrentMonth() {
	date := time.Now()
	fields.date.SetText(fmt.Sprint(getMonthInPortuguese(date.Month()), "/", date.Year()))
}

func (fields FinancialFields) subtractMonth() {
	_, date := fields.GetDate()
	dateNew := date.AddDate(0, -1, 0)
	if dateNew.Month() == date.Month() {dateNew = dateNew.AddDate(0, -1, 0)}
	finalDate := fmt.Sprint(getMonthInPortuguese(dateNew.Month()), "/", dateNew.Year())
	fields.date.SetText(finalDate)
	fields.date.ShowNow()
}

func (fields FinancialFields) addMonth() {
	_, date := fields.GetDate()
	dateNew := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0)
	finalDate := fmt.Sprint(getMonthInPortuguese(dateNew.Month()), "/", dateNew.Year())
	fields.date.SetText(finalDate)
	fields.date.ShowNow()
}