package ui

import (
	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

func addClientPage(db *gorm.DB, edit *gtk.Button, notebook *gtk.Notebook) *gtk.Box {
	thisPage := notebook.GetNPages()
	scrollable, err := gtk.ScrolledWindowNew(nil, nil)
	handler.Error("ui/Client.go >> scrollable, gtk.ScrolledWindow: ", err)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Client.go >> box", err)

	handlers, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Client.go >> handlers, gtk.Box: ", err)

	//Client Expander
	clientDataExpander, err := gtk.ExpanderNew("Dados Pessoais")
	handler.Error("ui/Client.go >> Line 33", err)
	
	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Client.go >> Line 35", err)

	//Fields
	n := widgets.PreForm("Nome", form)
	c := widgets.PreForm("CPF/CNPJ", form)
	r := widgets.PreForm("RG", form)
	nas := widgets.PreFormCalendar("Data de nascimento", form)
	s := widgets.PreFormComboBox("Sexo", []string{"Feminino", "Masculino", "Indefinido"}, form)
	tp := widgets.PreFormComboBox("Tipo de pessoa", []string {"Fisica", "Juridica"}, form)
	ec := widgets.PreFormComboBox("Estado Civil", []string {"solteiro", "casado", "separado", "divorciado", "viúvo"}, form)

	clientDataExpander.Add(form)
	handlers.PackStart(clientDataExpander, false, true, 10)

	//Address of the client
	clientAddressExpander, err := gtk.ExpanderNew("Informações para contato")
	handler.Error("ui/Client.go >> clientAddress, gtk.Expander: ", err)
	
	form2, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Client.go >> form2, gtk.Box: ", err)

	//Fields
	t := widgets.PreForm("Telefone", form2)
	w := widgets.PreFormCheckBox("WhatsApp", form2)
	ta := widgets.PreForm("Outro Telefone", form2)
	e := widgets.PreForm("E-mail", form2)
	pn := widgets.PreForm("Pais Natal", form2)
	en := widgets.PreForm("Estado Natal", form2)
	cn := widgets.PreForm("Cidade Natal", form2)

	notebook.Connect("switch-page", func (_ *gtk.Notebook, _ *gtk.Widget, index int)  {
		if index == thisPage {
			image, err := gtk.ImageNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
			handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
			edit.SetImage(image)
		} else {
			image, err := gtk.ImageNewFromIconName("document-edit-symbolic", gtk.ICON_SIZE_BUTTON)
			handler.Error("controller/ResultController.go >> notebook.Connect() >> image new from icon name", err)
			edit.SetImage(image)
		}
	})

	clientAddressExpander.Add(form2)
	handlers.PackStart(clientAddressExpander, false, true, 10)

	scrollable.Add(handlers)
	box.PackStart(scrollable, true, true, 0)
	controller.ClientInit(n, c, r, nas, s, tp, ec, t, w, ta, e, pn, en, cn, edit, db)
	return box
}