package ui

import (
	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

var addClient *gtk.Button
var addProject *gtk.Button
var editButton *gtk.Button

var OpenedAddClientPage = false
var OpenedAddProjectPage = false

func OnActivate(app *gtk.Application, db *gorm.DB) {
	loginWin, _ := gtk.ApplicationWindowNew(app)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	users, _ := gtk.ComboBoxTextNew()
	pass := widgets.PasswordEntry()
	button, _ := gtk.ButtonNewWithLabel("Entrar")
	loadButtonStyle(button)

	users.SetMarginStart(10)
	users.SetMarginEnd(10)
	pass.SetMarginStart(10)
	pass.SetMarginEnd(10)
	button.SetMarginStart(10)
	button.SetMarginEnd(10)

	loginWin.SetTitle("Entrar")
	loginWin.SetDefaultSize(300, 100)
	loginWin.SetPosition(gtk.WIN_POS_CENTER)
	box.PackStart(users, false, false, 5)
	box.PackStart(pass, false, false, 5)
	box.PackStart(button, false, false, 10)

	for _, u := range controller.GetUsers(db) {
		users.AppendText(u.Nome)
	}

	loginWin.Add(box)
	loginWin.ShowAll()

	button.Connect("clicked", func () {
		password, _ := pass.GetText()
		logged, employee := controller.Login(users.GetActiveText(), password, db)
		if !logged {return}
		println("Logged named as ", employee.Nome)
		preUi(app, loginWin, employee, db)
	})
}

func preUi(app *gtk.Application, loginwin *gtk.ApplicationWindow, employee database.Employees, db *gorm.DB) {
	win, _ := gtk.ApplicationWindowNew(app)
	win.SetDefaultSize(1200, 700)
	win.SetPosition(gtk.WIN_POS_CENTER)

	pixbuf, _ := gdk.PixbufNewFromFile("resources/logo.jpeg")
	win.SetIcon(pixbuf)

	bar, hbb, searchbar := getHeaderbar()

	addClient = hbb[0]
	addProject = hbb[1]
	editButton = hbb[2]

	win.SetTitlebar(bar)

	ui(win, searchbar, db, &employee)
	win.ShowAll()
	win.GrabFocus()
	loginwin.Hide()
}

func ui(win *gtk.ApplicationWindow, searchbar *gtk.SearchEntry, db *gorm.DB, user *database.Employees) {
	notebook, err := gtk.NotebookNew()
	handler.Error("ui/handler/FlexUi.go >> Line 19", err)

	controller.InitSearch(searchbar, notebook, db, editButton) 

	lbl1, _ := gtk.LabelNew("Inicio")
	homeBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0) 
	homeBox.PackStart(Home(db, editButton, notebook, user), true, true, 0)
	notebook.AppendPage(homeBox, lbl1)
	notebook.ShowAll()

	win.Add(notebook)

	addClient.Connect("clicked", func() {
		if OpenedAddClientPage {return}
		notebook.AppendPage(addClientPage(db, editButton, notebook), MakeTabs("Cadastrar Cliente", notebook))
		notebook.ShowAll()
		OpenedAddClientPage = true
	})

	addProject.Connect("clicked", func() {
		if OpenedAddProjectPage {return}
		notebook.AppendPage(addProjectPage(db, editButton, notebook), MakeTabs("Criar um Projeto", notebook))		
		notebook.ShowAll()
		OpenedAddProjectPage = true
	})
}

func loadButtonStyle(button *gtk.Button) {
	prov, _ := gtk.CssProviderNew()
	prov.LoadFromPath("resources/buttons.css")
	con, _ := button.GetStyleContext()
	con.AddProvider(prov, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

func MakeTabs(name string, notebook *gtk.Notebook) (*gtk.Box) {
	l, err := gtk.LabelNew(name)
	handler.Error("ui/FlexUi.go/ func Ui() >> maketabs >> label for " + name, err)
	
	close, err := gtk.ButtonNewWithLabel("✕")
	handler.Error("ui/FlexUi.go/ func Ui() >> maketabs >> button close for " + name, err)
	close.SetRelief(gtk.RELIEF_NONE)

	tab, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	handler.Error("ui/FlexUi.go/ func Ui() >> maketabs >> tab for " + name, err)
	
	tab.PackStart(l, true, true, 2)
	tab.PackEnd(close, false, true, 0)
	tab.ShowAll()

	close.Connect("clicked", func ()  {
		if name == "Cadastrar Cliente" {OpenedAddClientPage = false}
		if name == "Criar um Projeto" {OpenedAddProjectPage = false}
		current := notebook.GetCurrentPage()
		if current == 0 {return}
		notebook.RemovePage(current)
	})

	return tab
}