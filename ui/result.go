package ui

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

func Result() (*gtk.Box) {
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
	preForm("Nome", clientForm)
	preForm("CPF/CNPJ", clientForm)
	preForm("RG", clientForm)
	preForm("Data de nascimento", clientForm)
	preForm("Sexo", clientForm)
	preFormComboBox("Tipo de pessoa", []string {"Fisica", "Juridica"}, clientForm)
	preFormComboBox("Estado Civil", []string {"solteiro", "casado", "separado", "divorciado", "viúvo"}, clientForm)
	preForm("Telefone", clientForm)
	preForm("WhatsApp", clientForm)
	preForm("Outro Telefone", clientForm)
	preForm("E-mail", clientForm)
	preForm("Pais Natal", clientForm)
	preForm("Estado Natal", clientForm)
	preForm("Cidade Natal", clientForm)
	client.Add(clientForm)
	handlers.PackStart(client, false, false, 10)

	//Project
	handlers.PackStart(project(), false, false, 10)

	scrollable.Add(handlers)
	box.PackStart(scrollable, true, true, 0)
	return box
}

func project() *gtk.Expander {
	project, err := gtk.ExpanderNew("Projeto localizado em >> ")
	handler.Error("ui/Result.go >> client, gtk.BoxNew: ", err)
	ProjectForm, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Result.go >> clientForm, gtk.BoxNew: ", err)
	preForm("Cidade", ProjectForm)
	preForm("Estado", ProjectForm)
	preForm("Bairro", ProjectForm)
	preForm("Endereço", ProjectForm)
	preForm("Número", ProjectForm)
	preForm("Complemento", ProjectForm)
	preFormComboBox("Status do projeto", 
	[]string {"Inicial", "Pagamento Inicial Confirmado", "Em produção", "Instalado", "Pagamento Final Confirmado", "Finalizado"}, 
	ProjectForm)
	preFormTextView("Observações", ProjectForm)
	preForm("Valor do projeto", ProjectForm)
	preFormCheckBox("Projeto por contrato", ProjectForm)
	
	
	makeEnviroment := func () (*gtk.Expander) {
		expander, err := gtk.ExpanderNew("Ambiente")
		handler.Error("ui/Project.go >> addExpanderForEnviroment >> expander, gtk.Expander", err)
		
		form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
		handler.Error("ui/Project.go >> addExpanderForEnviroment >> form, gtk.Box", err)
	
		value := preForm("Nome do Ambiente", form)
		preFormTextView("Materiais", form)
		preFormCalendar("Data de fabricação", form)
		preFormCalendar("Data de instalação", form)
		
		value.Connect("changed", func() {
			s, _ := value.GetText()
			expander.SetLabel(s)
		})
		expander.Add(form)
		return expander
	}
	
	v := makeEnviroment()
	v.SetMarginStart(10)
	ProjectForm.PackEnd(v, false, false, 10)

	project.Add(ProjectForm)
	return project
}