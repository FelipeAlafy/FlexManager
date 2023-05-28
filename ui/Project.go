package ui

import (
	"strconv"
	"strings"

	"github.com/FelipeAlafy/Flex/controller"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/FelipeAlafy/Flex/widgets"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

const (
	COLUMN_PAYMENT_TYPE = iota
	COLUMN_VALUE
	COLUMN_OBSERVATION
)

func addProjectPage(db *gorm.DB, edit *gtk.Button, notebook *gtk.Notebook) (*gtk.Box) {
	thisPage := notebook.GetNPages()

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
	clients := widgets.PreFormComboBox("Cliente", []string {}, form) // This array need to be replaced with the data from database
	cep := widgets.PreForm("CEP", form)
	cidade := widgets.PreForm("Cidade", form)
	estado := widgets.PreForm("Estado", form)
	bairro := widgets.PreForm("Bairro", form)
	endereco := widgets.PreForm("Endereço", form)
	numero := widgets.PreForm("Número", form)
	complemento := widgets.PreForm("Complemento", form)
	status := widgets.PreFormComboBox("Status do projeto", 
	[]string {"Inicial", "Pagamento Inicial Confirmado", "Em produção", "Instalado", "Pagamento Final Confirmado", "Finalizado"}, 
	form)
	observacoes := widgets.PreFormTextView("Observações", form)
	contrato := widgets.PreFormCheckBox("Projeto por contrato", form)
	
	//Payment
	storage, value, payCombo, obs, add, vl  := widgets.PreFormForPay(form, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION)
	
	add.Connect("clicked", func ()  {
		paytype := payCombo.GetActiveText()
		s, _ := value.GetText()
		o, _ := obs.GetText()
		if paytype == "" || s == "" {return}
		widgets.AddRow(storage, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION, paytype, s, o, vl)
	})

	//End

	controller.InitProject(db, handlers, clients, cep, cidade, estado, bairro, endereco,
		 numero, complemento, status, observacoes, contrato, value, payCombo, obs, vl)
	
	//Funcionarios envolvidos, table with the name of the emploees
	addEnviroment, err := gtk.ButtonNewWithLabel("Adicionar um ambiente a este projeto")
	handler.Error("ui/Project.go >> addEnviroment, gtk.Button", err)
	form.PackEnd(addEnviroment, true, true, 10)

	//Variables
	Envs := []controller.EnvFields{}
	Expanders := []*gtk.Expander{}

	addEnviroment.Connect("clicked", func ()  {
		ex, env := addExpanderForEnviroment()
		Expanders = append(Expanders, ex)
		Envs = append(Envs, env)
		handlers.PackStart(ex, false, false, 10)
		handlers.ShowAll()
	})

	projectExpander.Add(form)
	handlers.PackStart(projectExpander, false, false, 10)
	projectExpander.SetExpanded(true)

	scrollable.Add(handlers)
	box.PackStart(scrollable, true, true, 0)

	edit.Connect("clicked", func ()  {
		controller.SaveProject(Envs, Expanders, storage)
	})

	return box
}

func addExpanderForEnviroment() (*gtk.Expander, controller.EnvFields) {
	expander, err := gtk.ExpanderNew("Ambiente")
	handler.Error("ui/Project.go >> addExpanderForEnviroment >> expander, gtk.Expander", err)
	
	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Project.go >> addExpanderForEnviroment >> form, gtk.Box", err)

	value := widgets.PreForm("Nome do Ambiente", form)
	materials := widgets.PreFormTextView("Materiais", form)
	production := widgets.PreFormCalendar("Data de fabricação", form)
	installation := widgets.PreFormCalendar("Data de instalação", form)

	env := controller.EnvFields{Name: value, Materials: materials, Production: production, Installation: installation}
	
	value.Connect("changed", func() {
		s, _ := value.GetText()
		expander.SetLabel(s)
	})

	expander.Add(form)
	return expander, env
}

func ToFloat(entry *gtk.Entry) float64 {
	entryValue, _ := entry.GetText()
	parser := strings.ReplaceAll(entryValue, ",", ".")
	v, err := strconv.ParseFloat(parser, 64)
	handler.Error("ui/widgets.go >> toFloat while trying to convert", err)
	return v
}