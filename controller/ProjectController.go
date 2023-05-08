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
	v *gtk.Entry,
	con *gtk.CheckButton) {
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
	valor = v
	contrato = con

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

func getModelForProject(envsBase []EnvFields) database.Project {
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
		ValorProjeto: toFloat(getDataFromEntry(valor)),
		Contrato: getDataFromCheckBox(contrato),
		Enviroments: getModelForEnviroments(envsBase),
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

func SaveProject(envsBase []EnvFields, expanders []*gtk.Expander) {
	model := getModelForProject(envsBase)
	name := getDataFromComboBox(clientBox)
	client := database.Client{Nome: name}
	cs := client.Search(dbProject)
	c := database.Client{}
	for _, v := range cs {
		c = v
	}
	c.AddProject(dbProject, model)
	clearProjectPage(expanders)
}

func getClients() []database.Client {
	c := database.GetAllClients(dbProject)
	println(c)
	return c
}

func clearProjectPage(expanders []*gtk.Expander) {
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
	for _, e := range expanders {
		handlers.Remove(e)
	}
}