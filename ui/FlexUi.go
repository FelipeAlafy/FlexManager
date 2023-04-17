package ui

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

var addClient *gtk.Button
var addProject *gtk.Button
var editButton *gtk.Button

var openedAddClientPage = false
var openedAddProjectPage = false

func OnActivate(app *gtk.Application) {
	win, err := gtk.ApplicationWindowNew(app)
	handler.Error("ui/handler/FlexUi.go >> Line 11", err)

	win.SetDefaultSize(900, 700)
	win.SetPosition(gtk.WIN_POS_CENTER)
	bar, hbb, searchbar := getHeaderbar()

	addClient = hbb[0]
	addProject = hbb[1]
	editButton = hbb[2]

	win.SetTitlebar(bar)

	ui(win, searchbar)
	win.ShowAll()
}

func ui(win *gtk.ApplicationWindow, searchbar *gtk.SearchEntry) {
	notebook, err := gtk.NotebookNew()
	handler.Error("ui/handler/FlexUi.go >> Line 19", err)

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
			if name == "Cadastrar Cliente" {openedAddClientPage = false}
			if name == "Criar um Projeto" {openedAddProjectPage = false}
			notebook.RemovePage(notebook.GetCurrentPage())
		})

		return tab
	} 

	lbl1, _ := gtk.LabelNew("Inicio")
	con, _ := gtk.ButtonNewWithLabel("This is an example of page")

	notebook.AppendPage(con, lbl1)
	notebook.GetCurrentPage()

	win.Add(notebook)

	addClient.Connect("clicked", func() {
		if openedAddClientPage {return}
		notebook.AppendPage(addClientPage(), maketabs("Cadastrar Cliente", notebook))
		notebook.ShowAll()
		openedAddClientPage = true
	})

	addProject.Connect("clicked", func() {
		if openedAddProjectPage {return}
		notebook.AppendPage(addProjectPage(), maketabs("Criar um Projeto", notebook))		
		notebook.ShowAll()
		openedAddProjectPage = true
	})

	editButton.Connect("clicked", func ()  {
		notebook.AppendPage(Result(), maketabs("Resultado", notebook))
		notebook.ShowAll()
	})

	searchbar.Connect("activate", func ()  {
		v, _ := searchbar.GetText()
		if len(v) > 2 {
			println("Starting the searching for clients and projects >> ", v)
			popup(searchbar)
		}
	})
}

