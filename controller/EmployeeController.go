package controller

import (
	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"

	"github.com/jinzhu/gorm"
)



type EmployeeFields struct {
	DB 				*gorm.DB
	CombotextName 	*gtk.ComboBoxText
	Password 		*gtk.Entry
	PasswordChecker *gtk.Entry
	checker 		*gtk.Image
	phone 			*gtk.Entry
	level 			*gtk.ComboBoxText
	editButton		*gtk.Button
}

func EmployeeInit(db *gorm.DB, combonome *gtk.ComboBoxText, password, passwordChecker *gtk.Entry, checker *gtk.Image, phone *gtk.Entry, level *gtk.ComboBoxText, editButton *gtk.Button) {
	fields := EmployeeFields{db, combonome, password, passwordChecker, checker, phone, level, editButton}
	thisPage := notebook.GetNPages()

	notebook.Connect("page-removed", func (_ *gtk.Notebook, _ *gtk.Widget, pageRemoved uint)  {
		if pageRemoved < uint(thisPage) {thisPage = thisPage - 1}
	})

	notebook.Connect("switch-page", func (_ *gtk.Notebook, _ *gtk.Widget, index int)  {
		if thisPage != index {return}
		image, err := gtk.ImageNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
		handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
		fields.editButton.SetImage(image)
	})

	fields.Password.Connect("changed", func ()  {
		println("both password have matched -> ", fields.ValidatePassword())
	})
	fields.PasswordChecker.Connect("changed", func ()  {
		println("both password have matched -> ", fields.ValidatePassword())
	})
	fields.checker.SetVisible(false)

	for _, u := range GetUsers(fields.DB) {
		fields.CombotextName.AppendText(u.Nome)
	}

	fields.CombotextName.Connect("changed", func ()  {
		for _, u := range GetUsers(fields.DB) {
			if u.Nome != fields.CombotextName.GetActiveText() {continue}
			fields.loadData(&u)
		}
	})

	fields.editButton.Connect("clicked", func ()  {
		fields.save()
		fields.clear()
		for _, u := range GetUsers(fields.DB) {
			fields.CombotextName.AppendText(u.Nome)
		}
	})
}

func (fields EmployeeFields) ValidatePassword() bool {
	pass, _ := fields.Password.GetText()
	passChecker, _ := fields.PasswordChecker.GetText()
	if  pass == passChecker && pass != "" {
		fields.checker.SetVisible(true)
		return true
	}
	
	fields.checker.SetVisible(false)
	return false
}

func (fields EmployeeFields) loadData(Employee *database.Employees) {
	fields.Password.SetText(Employee.Senha)
	fields.PasswordChecker.SetText(Employee.Senha)
	fields.phone.SetText(Employee.Telefone)
	fields.checker.SetSensitive(true)
}

func GetUsers(dbUsers *gorm.DB) []database.Employees {
	return database.GetAllEmployees(dbUsers)
}

func (fields EmployeeFields) getModelForEmployee() database.Employees {
	name := fields.CombotextName.GetActiveText()
	pass, _ := fields.Password.GetText()
	phone, _ := fields.phone.GetText()
	level:= fields.level.GetActiveText()
	levelNumber := uint(1)
	switch level {
	case "Padrão":
		levelNumber = 1
	case "Administrativo":
		levelNumber = 2
	default:
		levelNumber = 1
	}
	return database.Employees{Nome: name, Telefone: phone, Senha: pass, Level: levelNumber}
}

func (fields EmployeeFields) save() {
	model := fields.getModelForEmployee()
	model.Save(fields.DB)
}

func (fields EmployeeFields) clear() {
	fields.checker.SetVisible(false)
	fields.CombotextName.SetActive(-1)
	fields.CombotextName.RemoveAll()
	entry, err := fields.CombotextName.GetEntry()
	handler.MinorError("controller/EmployeeController.go >> clear >> entry", err)
	entry.SetText("")
	fields.phone.SetText("")
	fields.Password.SetText("")
	fields.PasswordChecker.SetText("")
	fields.level.SetActive(0)
}