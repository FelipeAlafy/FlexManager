package controller

import (
	"fmt"
	"time"

	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

var HomeDB *gorm.DB
var editButton *gtk.Button
var notebook *gtk.Notebook

type Graph struct {
	GraphType uint
	Values []float64
}

func HomeInit(homedb *gorm.DB, mainBox, projects *gtk.Box, edit *gtk.Button, note *gtk.Notebook) {
	HomeDB = homedb
	editButton = edit
	notebook = note
	thisPage := 0

	notebook.Connect("switch-page", func (_ *gtk.Notebook, _ *gtk.Widget, index int)  {
		if thisPage != index {return}
		image, err := gtk.ImageNewFromIconName("document-edit-symbolic", gtk.ICON_SIZE_BUTTON)
		handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
		editButton.SetImage(image)
	})


	populate(getProjects(), projects)
}

func getProjects() []database.Project {
	return database.GetProjectsByLastEdit(HomeDB)
}

func populate(projects []database.Project, projectsBox *gtk.Box) {
	c := 0
	for _, p := range projects {
		str := fmt.Sprintf("Cidade: %s, %s\nEndereço: %s\nNúmero: %d\nÚltima Edição: %s", p.Cidade, p.Bairro, p.Endereco, p.Numero, getDate(p.UpdatedAt))
		if p.Complemento != "" {
			str = fmt.Sprintf("Cidade: %s, %s\nEndereço: %s\nNúmero: %d, %s\nÚltima Edição: %s", p.Cidade, p.Bairro, p.Endereco, p.Numero, p.Complemento, getDate(p.UpdatedAt))
		}

		frame, err := gtk.FrameNew("")
		handler.Error("controller/home.go >> populate >> frame", err)

		card, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		handler.Error("controller/Home.go >> populate >> card", err)

		btn, err := gtk.ButtonNewFromIconName("go-next-symbolic", gtk.ICON_SIZE_BUTTON)
		handler.Error("controller/Home.go >> populate >> btn", err)
		btn.SetRelief(gtk.RELIEF_NONE)

		btn.Connect("clicked", func ()  {
			for _, pc := range projects {
				strc := fmt.Sprintf("Cidade: %s, %s\nEndereço: %s\nNúmero: %d\nÚltima Edição: %s", pc.Cidade, pc.Bairro, pc.Endereco, pc.Numero, getDate(pc.UpdatedAt))
				if pc.Complemento != "" {
					strc = fmt.Sprintf("Cidade: %s, %s\nEndereço: %s\nNúmero: %d, %s\nÚltima Edição: %s", pc.Cidade, pc.Bairro, pc.Endereco, pc.Numero, pc.Complemento, getDate(pc.UpdatedAt))
				}
				if strc != str {
					continue
				}
				client := database.Search(HomeDB, pc.ClientId)
				makeTabForResult(fmt.Sprintf("Resultados para %s", client.Nome), notebook, client, editButton)
				notebook.ShowAll()
			}
		})

		lbl, err := gtk.LabelNew(str)
		handler.Error("controller/home.go >> populate >> lbl", err)

		
		card.PackStart(lbl, true, true, 0)
		card.PackEnd(btn, false, false, 5)
		frame.Add(card)

		FrameProv, _ := gtk.CssProviderNew()
		FrameProv.LoadFromPath("resources/frame.css")
		FrameContext, _ := frame.GetStyleContext()
		FrameContext.AddProvider(FrameProv, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		ButtonProv, _ := gtk.CssProviderNew() 
		ButtonProv.LoadFromPath("resources/buttons.css")
		ButtonContext, _ := btn.GetStyleContext()
		ButtonContext.AddProvider(ButtonProv, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		projectsBox.PackStart(frame, true, true, 10)
		c++
		projectsBox.ShowAll()
	}
}

func getDate(time time.Time) string {
	return fmt.Sprintf("%d/%d/%d", time.Day(), time.Month(), time.Year())	
}