package controller

import (
	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

func Result(c database.Client, db *gorm.DB, edit *gtk.Button, notebook *gtk.Notebook) *gtk.Box {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	handler.Error("ui/Project.go >> box, gtk.BoxNew: ", err)
	
	scrollable, err := gtk.ScrolledWindowNew(nil, nil)
	handler.Error("ui/Project.go >> scrollable, gtk.ScrolledWindow: ", err)

	handlers, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Project.go >> handlers, gtk.BoxNew: ", err)


	client, err := gtk.ExpanderNew("Dados de Cliente")
	handler.Error("ui/Result.go >> client, gtk.BoxNew: ", err)
	clientForm, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Result.go >> clientForm, gtk.BoxNew: ", err)
	nome := widgets.PreForm("Nome", clientForm)
	cpf := widgets.PreForm("CPF/CNPJ", clientForm)
	rg := widgets.PreForm("RG", clientForm)
	nascimento := widgets.PreFormCalendar("Data de nascimento", clientForm)
	sexo := widgets.PreFormComboBox("Sexo", []string {"Feminino", "Masculino", "Indefinido"}, clientForm)
	tp := widgets.PreFormComboBox("Tipo de pessoa", []string {"Fisica", "Juridica"}, clientForm)
	ec := widgets.PreFormComboBox("Estado Civil", []string {"solteiro", "casado", "separado", "divorciado", "viúvo"}, clientForm)
	telefone, whatsapp := widgets.PreFormContactWhatsapp("Telefone", clientForm)
	w := widgets.PreFormCheckBox("WhatsApp", clientForm)
	ta := widgets.PreForm("Outro Telefone", clientForm)
	email := widgets.PreForm("E-mail", clientForm)
	pn := widgets.PreForm("Pais Natal", clientForm)
	en := widgets.PreForm("Estado Natal", clientForm)
	cn := widgets.PreForm("Cidade Natal", clientForm)
	client.Add(clientForm)
	handlers.PackStart(client, true, true, 10)

	scrollable.Add(handlers)
	box.PackStart(scrollable, true, true, 0)

	//Client
	ClientField := ClientFields{
		Nome:        nome,
		CPF:         cpf,
		RG:          rg,
		Nascimento:  nascimento,
		Sexo:        sexo,
		TipoPessoa:  tp,
		EstadoCivil: ec,
		Telefone:    telefone,
		WhatsApp:    w,
		WhatsAppBtn: whatsapp,
		TelefoneAlt: ta,
		Email:       email,
		PaisNatal:   pn,
		EstadoNatal: en,
		CidadeNatal: cn,
		Project:     make([]ProjectFields, len(c.Projects)),
	}

	InitResult(ClientField, c, handlers, db, edit, notebook)

	return box
}