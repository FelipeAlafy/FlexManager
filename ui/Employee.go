package ui

import (
	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

func Employee(notebook *gtk.Notebook, db*gorm.DB, edit *gtk.Button) *gtk.Box {
	MainBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Employee.go >> Employee >> MainBox", err)

	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Employee.go >> Employee >> form", err)

	combonome := widgets.PreFormComboBoxWithAnEntry("Nome", []string{}, form)
	password := widgets.PreFormPassword("Senha", form)
	passwordChecker, checker := widgets.PreFormPasswordEntryWithConfirmation("Senha", form)
	phone := widgets.PreForm("Telefone", form)
	level := widgets.PreFormComboBox("Nível de acesso", []string{"Padrão", "Administrativo"}, form)

	controller.EmployeeInit(db, combonome, password, passwordChecker, checker, phone, level, edit)

	MainBox.PackStart(form, true, false, 10)
	checker.SetVisible(false)
	return MainBox
}