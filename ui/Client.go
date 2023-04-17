package ui

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

func addClientPage() (*gtk.Box) {
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
	preForm("Nome", form)
	preForm("CPF/CNPJ", form)
	preForm("RG", form)
	preForm("Data de nascimento", form)
	preForm("Sexo", form)
	preFormComboBox("Tipo de pessoa", []string {"Fisica", "Juridica"}, form)
	preFormComboBox("Estado Civil", []string {"solteiro", "casado", "separado", "divorciado", "viúvo"}, form)

	clientDataExpander.Add(form)
	handlers.PackStart(clientDataExpander, false, true, 10)

	//Address of the client
	clientAddressExpander, err := gtk.ExpanderNew("Informações para contato")
	handler.Error("ui/Client.go >> clientAddress, gtk.Expander: ", err)
	
	form2, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Client.go >> form2, gtk.Box: ", err)

	//Fields
	preForm("Telefone", form2)
	preForm("WhatsApp", form2)
	preForm("Outro Telefone", form2)
	preForm("E-mail", form2)
	preForm("Pais Natal", form2)
	preForm("Estado Natal", form2)
	preForm("Cidade Natal", form2)

	clientAddressExpander.Add(form2)
	handlers.PackStart(clientAddressExpander, false, true, 10)


	scrollable.Add(handlers)
	box.PackStart(scrollable, true, true, 0)
	return box
}