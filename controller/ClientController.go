package controller

import (
	"time"

	"github.com/FelipeAlafy/Flex/database"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jinzhu/gorm"
)

//Fields
var nome *gtk.Entry
var cpf *gtk.Entry
var rg *gtk.Entry
var nascimento *gtk.Calendar
var sexo *gtk.ComboBoxText
var tipoPessoa *gtk.ComboBoxText
var estadoCivil *gtk.ComboBoxText
var telefone *gtk.Entry
var whatsapp *gtk.CheckButton
var telefoneAlt *gtk.Entry
var email *gtk.Entry
var paisNatal *gtk.Entry
var estadoNatal *gtk.Entry
var cidadeNatal *gtk.Entry
var edit *gtk.Button
var db *gorm.DB

func ClientInit(n *gtk.Entry, c *gtk.Entry, r *gtk.Entry, nas *gtk.Calendar, s *gtk.ComboBoxText,
	tp *gtk.ComboBoxText, ec *gtk.ComboBoxText, t *gtk.Entry, w *gtk.CheckButton, ta *gtk.Entry,
	e *gtk.Entry, pn *gtk.Entry, en *gtk.Entry, cn*gtk.Entry, button *gtk.Button,
	database *gorm.DB) {
	nome = n
	cpf = c
	rg = r
	nascimento = nas
	sexo = s
	tipoPessoa = tp
	estadoCivil = ec
	telefone = t
	whatsapp = w
	telefoneAlt = ta
	email = e
	paisNatal = pn
	estadoNatal = en
	cidadeNatal = cn
	edit = button
	db = database

	edit.Connect("clicked", func () {
		println("Trying to save a client into database")
		model := getModelForClient()
		SaveClient(model)
	})
}

func SaveClient(model database.Client) {
	model.New(db)
	Clear()
}

func getModelForClient() database.Client {
	return database.Client{
		Nome: getDataFromEntry(nome),
		Cpf: getDataFromEntry(cpf),
		Rg: getDataFromEntry(rg),
		Nascimento: getDataFromCalendar(nascimento),
		Sexo: getDataFromComboBox(sexo),
		TipoPessoa: getDataFromComboBox(tipoPessoa),
		EstadoCivil: getDataFromComboBox(estadoCivil),
		Telefone: getDataFromEntry(telefone),
		Whatsapp: getDataFromCheckBox(whatsapp),
		TelefoneAlt: getDataFromEntry(telefoneAlt),
		Email: getDataFromEntry(email),
		PaisNatal: getDataFromEntry(paisNatal),
		EstadoNatal: getDataFromEntry(estadoNatal),
		CidadeNatal: getDataFromEntry(cidadeNatal),
	}
}

func Clear() {
	time := time.Now()
	nome.SetText("")
	cpf.SetText("")
	rg.SetText("")
	nascimento.SelectDay(uint(time.Day()))
	nascimento.SelectMonth(uint(time.Month()), uint(time.Year()))
	sexo.SetActive(-1)
	tipoPessoa.SetActive(-1)
	estadoCivil.SetActive(-1)
	telefone.SetText("")
	whatsapp.SetActive(false)
	telefoneAlt.SetText("")
	email.SetText("")
	paisNatal.SetText("")
	estadoNatal.SetText("")
	cidadeNatal.SetText("")
}