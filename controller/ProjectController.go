package controller

import (
	"github.com/FelipeAlafy/Flex/database"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

type EnvFields struct {
	Name 			*gtk.Entry
	Materials 		*gtk.TextView
	Production 		*gtk.Calendar
	Installation 	*gtk.Calendar
}

type PayFields struct {
	Value 			*gtk.Entry
	Way 			*gtk.ComboBoxText
	Observacoes 	*gtk.TextView
}

var dbProject 		*gorm.DB
var clientBox 		*gtk.ComboBoxText
var cep 			*gtk.Entry
var cidade 			*gtk.Entry
var estado 			*gtk.Entry
var bairro 			*gtk.Entry
var endereco 		*gtk.Entry
var numero 			*gtk.Entry
var complemento 	*gtk.Entry
var status 			*gtk.ComboBoxText
var observacoes 	*gtk.TextView
var valor 			*gtk.Entry
var contrato 		*gtk.CheckButton
var handlers		*gtk.Box
var payValue		*gtk.Entry
var payCombo		*gtk.ComboBoxText
var payObservation	*gtk.Entry
var ValLabel		*gtk.Label

func InitProject(
	gormDB *gorm.DB,
	hand *gtk.Box,
	cli *gtk.ComboBoxText,
	c *gtk.Entry,
	ci *gtk.Entry,
	e *gtk.Entry,
	b *gtk.Entry,
	en *gtk.Entry,
	n *gtk.Entry,
	co *gtk.Entry,
	s *gtk.ComboBoxText,
	o *gtk.TextView,
	con *gtk.CheckButton,
	payVal *gtk.Entry,
	payC *gtk.ComboBoxText,	
	payObs *gtk.Entry,
	vl		*gtk.Label,
	) {
	dbProject = gormDB
	handlers = hand
	clientBox = cli
	cep = c
	cidade = ci
	estado = e
	bairro = b
	endereco = en
	numero = n
	complemento = co
	status = s
	observacoes = o
	contrato = con
	payValue = payVal
	payCombo = payC
	payObservation = payObs
	ValLabel = vl
	thisPage := notebook.GetNPages()

	notebook.Connect("page-removed", func (_ *gtk.Notebook, _ *gtk.Widget, pageRemoved uint)  {
		if pageRemoved < uint(thisPage) {thisPage = thisPage - 1}
	})

	notebook.Connect("switch-page", func (_ *gtk.Notebook, _ *gtk.Widget, index int)  {
		if thisPage != index {return}
		image, err := gtk.ImageNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
		handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
		editButton.SetImage(image)
	})

	SyncClientComboBox()

	cep.Connect("activate", func ()  {
		address, err := cep.GetText()
		handler.Error("controller/ClientController.go >> address, cep.GetText()", err)
		cepMap := handler.GetAddress(address)
		cidade.SetText(cepMap.Localidade)
		bairro.SetText(cepMap.Bairro)
		endereco.SetText(cepMap.Logradouro)
		estado.SetText(cepMap.UF)
	})
}

func SyncClientComboBox() {
	clientBox.RemoveAll()
	clients := getClients()
	for _, client := range clients {
		clientBox.AppendText(client.Nome)
	}
}

func getModelForProject(envsBase []EnvFields, listStore *gtk.ListStore) database.Project {
	model := database.Project{
		CEP: getDataFromEntry(cep),
		Cidade: getDataFromEntry(cidade),
		Estado: getDataFromEntry(estado),
		Bairro: getDataFromEntry(bairro),
		Endereco: getDataFromEntry(endereco),
		Numero: uint(toInt(getDataFromEntry(numero))),
		Complemento: getDataFromEntry(complemento),
		Status: getDataFromComboBox(status),
		Observacoes: getDataFromTextView(observacoes),
		Contrato: getDataFromCheckBox(contrato),
		Enviroments: getModelForEnviroments(envsBase),
		Payments: getModelForPayments(listStore),
	}
	return model
}

func getModelForEnviroments(envs []EnvFields) []database.Enviroment {
	enviroments := make([]database.Enviroment, len(envs))

	for i, env := range envs {
		enviroments[i] = database.Enviroment{
			Name: getDataFromEntry(env.Name),
			Materials: getDataFromTextView(env.Materials),
			Installation: getDataFromCalendar(env.Installation),
			Production: getDataFromCalendar(env.Production),
		}
	}

	return enviroments
}

func getModelForPayments(store *gtk.ListStore) []database.Payment {
	payments := []database.Payment{}

	store.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		payType, err := model.GetValue(iter, 0)
		if err != nil {return false}
		value, err := model.GetValue(iter, 1)
		if err != nil {return false}
		obs, err := model.GetValue(iter, 2)
		if err != nil {return false}

		v, err := value.GetString()
		if err != nil {return false}
		p, err := payType.GetString()
		if err != nil {return false}
		o, err := obs.GetString()
		if err != nil {return false}

		payment := database.Payment{
			Value: handler.ConvertStringIntoFloat(v),
			Way: p,
			Observation: o,
		}
		payments = append(payments, payment)

		if v == "" && p == "" {
			return true
		}

		return false
	})

	return payments
}

func SaveProject(envsBase []EnvFields, expanders []*gtk.Expander, listStorage *gtk.ListStore) {
	model := getModelForProject(envsBase, listStorage)
	name := getDataFromComboBox(clientBox)
	client := database.Client{Nome: name}
	cs := client.Search(dbProject)
	c := database.Client{}
	for _, v := range cs {
		c = v
	}
	c.AddProject(dbProject, model)
	clearProjectPage(expanders, listStorage)
}

func getClients() []database.Client {
	c := database.GetAllClients(dbProject)
	println(c)
	return c
}

func clearProjectPage(expanders []*gtk.Expander, storage *gtk.ListStore) {
	clientBox.SetActive(-1)
	cep.SetText("")
	cidade.SetText("")
	estado.SetText("")
	bairro.SetText("")
	endereco.SetText("")
	numero.SetText("")
	complemento.SetText("")
	status.SetActive(-1)
	b, _ := observacoes.GetBuffer()
	b.SetText("")
	valor.SetText("")
	contrato.SetActive(false)
	payCombo.SetActive(-1)
	payValue.SetText("")
	payObservation.SetText("")
	ValLabel.SetText("0,00")
	for _, e := range expanders {
		handlers.Remove(e)
	}

	storage.Clear()
}