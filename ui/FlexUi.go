package ui

import (
	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

var addClient *gtk.Button
var addProject *gtk.Button
var editButton *gtk.Button

var OpenedAddClientPage = false
var OpenedAddProjectPage = false

func OnActivate(app *gtk.Application, db *gorm.DB) {
	win, err := gtk.ApplicationWindowNew(app)
	handler.Error("ui/handler/FlexUi.go >> Line 11", err)

	win.SetDefaultSize(900, 700)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetIconFromFile("union_logo.jpg")

	bar, hbb, searchbar := getHeaderbar()

	addClient = hbb[0]
	addProject = hbb[1]
	editButton = hbb[2]

	win.SetTitlebar(bar)

	ui(win, searchbar, db)
	win.ShowAll()
}

func ui(win *gtk.ApplicationWindow, searchbar *gtk.SearchEntry, db *gorm.DB) {
	notebook, err := gtk.NotebookNew()
	handler.Error("ui/handler/FlexUi.go >> Line 19", err)

	controller.InitSearch(searchbar, notebook, db, editButton)

	maketabs := func (name string, notebook *gtk.Notebook) (*gtk.Box) {
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

	lbl1, _ := gtk.LabelNew("Inicio")
	con, _ := gtk.ButtonNewWithLabel("This is an example of page")

	notebook.AppendPage(con, lbl1)
	notebook.GetCurrentPage()

	win.Add(notebook)

	addClient.Connect("clicked", func() {
		if OpenedAddClientPage {return}
		notebook.AppendPage(addClientPage(db), maketabs("Cadastrar Cliente", notebook))
		notebook.ShowAll()
		OpenedAddClientPage = true
	})

	addProject.Connect("clicked", func() {
		if OpenedAddProjectPage {return}
		notebook.AppendPage(addProjectPage(db), maketabs("Criar um Projeto", notebook))		
		notebook.ShowAll()
		OpenedAddProjectPage = true
	})
}

