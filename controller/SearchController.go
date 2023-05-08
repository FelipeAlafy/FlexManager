package controller

import (
	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

var searchbar *gtk.SearchEntry
var dbSearch *gorm.DB

func InitSearch(Searchbar *gtk.SearchEntry, notebook *gtk.Notebook, dbs *gorm.DB, edit *gtk.Button) {
	searchbar = Searchbar
	dbSearch = dbs

	searchbar.Connect("activate", func ()  {
		v, _ := searchbar.GetText()
		if len(v) > 2 {
			box := Search(v, notebook, edit)
			popup(searchbar, box)
			
		}
	})
}

func Search(name string, notebook *gtk.Notebook, edit *gtk.Button) *gtk.Box {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/SearchPopup.go >> box", err)

	n := database.Client{Nome: name}
	clients := n.Search(dbSearch)

	for _, c := range clients {
		mainBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
		lbl, _ := gtk.LabelNew(c.Nome + "\nNasceu em: " + c.CidadeNatal)
		btn, _ := gtk.ButtonNewFromIconName("go-next-symbolic", gtk.ICON_SIZE_BUTTON)
		btn.SetRelief(gtk.RELIEF_NONE)
		mainBox.PackStart(lbl, true, true, 5)
		mainBox.PackEnd(btn, false, false, 5)
		box.PackStart(mainBox, false, true, 5)

		btn.Connect("clicked", func ()  {
			makeTabForResult("Resultados para " + c.Nome, notebook, c, edit)
		})
	}
	return box
}

func popup(root *gtk.SearchEntry, box *gtk.Box) {	
	pop, err := gtk.PopoverNew(root)
	handler.Error("ui/SearchPopup.go >> pop", err)
	pop.SetPosition(gtk.POS_BOTTOM)
	pop.Add(box)
	pop.ShowAll()
}

func makeTabForResult(name string, notebook *gtk.Notebook, c database.Client, edit *gtk.Button) {
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
		current := notebook.GetCurrentPage()
		if current == 0 {return}
		notebook.RemovePage(current)
	})

	notebook.AppendPage(Result(c, dbSearch, edit, notebook), tab)
	notebook.ShowAll()
}