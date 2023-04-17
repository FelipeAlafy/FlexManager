package ui

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

func addProjectPage() (*gtk.Box) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	handler.Error("ui/Project.go >> box, gtk.BoxNew: ", err)
	
	scrollable, err := gtk.ScrolledWindowNew(nil, nil)
	handler.Error("ui/Project.go >> scrollable, gtk.ScrolledWindow: ", err)

	handlers, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Project.go >> handlers, gtk.BoxNew: ", err)

	projectExpander, err :=  gtk.ExpanderNew("Dados do projeto")
	handler.Error("ui/Project.go >> projectExpander, gtk.projectExpander: ", err)

	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Project.go >> form, gtk.BoxNew: ", err)

	//Fields
	preFormComboBoxWithAnEntry("Cliente", 
	[]string {"Felipe", "Alafy", "Rodrigues", "Silva"}, form) // This array need to be replaced with the data from database
	preForm("Cidade", form)
	preForm("Estado", form)
	preForm("Bairro", form)
	preForm("Endereço", form)
	preForm("Número", form)
	preForm("Complemento", form)
	preFormComboBox("Status do projeto", 
	[]string {"Inicial", "Pagamento Inicial Confirmado", "Em produção", "Instalado", "Pagamento Final Confirmado", "Finalizado"}, 
	form)
	preFormTextView("Observações", form)
	preForm("Valor do projeto", form)
	preFormCheckBox("Projeto por contrato", form)
	
	//Funcionarios envolvidos, table with the name of the emploees

	addEnviroment, err := gtk.ButtonNewWithLabel("Adicionar um ambiente a este projeto")
	handler.Error("ui/Project.go >> addEnviroment, gtk.Button", err)
	form.PackStart(addEnviroment, true, true, 10)

	addEnviroment.Connect("clicked", func ()  {
		handlers.PackStart(addExpanderForEnviroment(), false, false, 10)
		handlers.ShowAll()
	})

	projectExpander.Add(form)
	handlers.PackStart(projectExpander, false, false, 10)

	scrollable.Add(handlers)
	box.PackStart(scrollable, true, true, 0)
	return box
}

func addExpanderForEnviroment() (*gtk.Expander) {
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