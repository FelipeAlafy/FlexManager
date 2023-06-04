package ui

import (
	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

func Home(homedb *gorm.DB, edit *gtk.Button, notebook *gtk.Notebook) *gtk.Box {
	homeBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	handler.Error("ui/home.go >> home >> homebox", err)
	
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	handler.Error("ui/home.go >> home >> scrolledWindow", err)
	
	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/home.go >> home >> form", err)

	expander, err := gtk.ExpanderNew("Projetos  → recentes")
	handler.Error("ui/home.go >> home >> expander for projects", err)

	projectBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	handler.Error("ui/home.go >> home >> projectBox", err)
	projectBox.SetHomogeneous(true)

	projects, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	handler.Error("ui/home.go >> home >> projects", err)

	centerBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handler.Error("ui/Home.go >> Home >> centerBox", err)

	centerBox.PackStart(createCenterButton("Ajuda", "resources/help.svg"), false, false, 30)
	centerBox.PackStart(createCenterButton("Financeiro", "resources/wallet.svg"), false, false, 30)
	centerBox.PackStart(createCenterButton("Empregados", "resources/employees.svg"), false, false, 30)
	centerBox.SetHAlign(gtk.ALIGN_CENTER)
	
	scrolledProjects, _ := gtk.ScrolledWindowNew(nil, nil)
	scrolledProjects.SetMinContentHeight(100)
	projectBox.PackStart(projects, true, true, 0)
	scrolledProjects.Add(projectBox)
	expander.Add(scrolledProjects)

	controller.HomeInit(homedb, homeBox, projects, edit, notebook)

	form.PackStart(expander, false, true, 0)
	form.PackStart(centerBox, false, false, 30)
	scrolledWindow.Add(form)
	homeBox.PackStart(scrolledWindow, true, true, 0)
	expander.SetExpanded(true)
	homeBox.ShowAll()
	return homeBox
}

func loadFrameStyle(frame *gtk.Frame) {
	prov, _ := gtk.CssProviderNew()
	prov.LoadFromPath("resources/frame.css")
	con, _ := frame.GetStyleContext()
	con.AddProvider(prov, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

func createCenterButton(label, path string) *gtk.Frame {
	frame, _ := gtk.FrameNew("")
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	btn, _ := gtk.ButtonNew()
	btn.SetRelief(gtk.RELIEF_NONE)

	pixbuf, _ := gdk.PixbufNewFromFileAtScale(path, 100, 100, true)
	img, _ := gtk.ImageNewFromPixbuf(pixbuf)
	btn.SetImage(img)

	helpLabel, _ := gtk.LabelNew(label)
	box.PackStart(btn, true, true, 5)
	box.PackStart(helpLabel, false, false, 0)
	frame.Add(box)
	loadFrameStyle(frame)
	return frame
}