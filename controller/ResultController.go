package controller

import (
	"fmt"
	"strings"

	"github.com/FelipeAlafy/Flex/database"
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

type ClientFields struct {
	Nome 			*gtk.Entry
	CPF 			*gtk.Entry
	RG 				*gtk.Entry
	Nascimento 		*gtk.Calendar
	Sexo 			*gtk.ComboBoxText
	TipoPessoa 		*gtk.ComboBoxText
	EstadoCivil 	*gtk.ComboBoxText
	Telefone 		*gtk.Entry
	WhatsApp 		*gtk.CheckButton
	WhatsAppBtn		*gtk.Button
	TelefoneAlt 	*gtk.Entry
	Email 			*gtk.Entry
	PaisNatal 		*gtk.Entry
	EstadoNatal 	*gtk.Entry
	CidadeNatal 	*gtk.Entry
	Project			[]ProjectFields
}

type ProjectFields struct {
	Cep 			*gtk.Entry
	Cidade 			*gtk.Entry
	Estado 			*gtk.Entry
	Bairro			*gtk.Entry
	Endereco 		*gtk.Entry
	Numero 			*gtk.Entry
	Complemento 	*gtk.Entry
	Status 			*gtk.ComboBoxText
	Observacoes 	*gtk.TextView
	Valor 			*gtk.Entry
	Contrato 		*gtk.CheckButton
	Payment 		PaymentFields
	Enviroments		[]EnviromentFields
}

type PaymentFields struct {
	PayCombo *gtk.ComboBoxText
	ValueEntry *gtk.Entry
	ObsEntry *gtk.Entry
	AddButton *gtk.Button
	Store		*gtk.ListStore
}

type EnviromentFields struct {
	Nome 			*gtk.Entry
	Materiais 		*gtk.TextView
	Fabricacao 		*gtk.Calendar
	Instalacao 		*gtk.Calendar
}


func InitResult(f ClientFields, c database.Client, handlers *gtk.Box, dbResult *gorm.DB, edit *gtk.Button, notebook *gtk.Notebook) {
	//These two variables control the flow of when edit button has a specific action or another one
	thisPage := notebook.GetNPages()
	buttonState := false

	f.Nome.SetText(c.Nome)
	f.CPF.SetText(c.Cpf)
	f.RG.SetText(c.Rg)
	f.Nascimento.SelectDay(uint(c.Nascimento.Day()))
	f.Nascimento.SelectMonth(uint(c.Nascimento.Month()), uint(c.Nascimento.Year()))
	
	switch (c.Sexo) {
	case "Feminino":
		f.Sexo.SetActive(0)
	case "Masculino":
		f.Sexo.SetActive(1)
	case "Indefinido":
		f.Sexo.SetActive(2)
	}
	
	switch (c.TipoPessoa) {
	case "Fisica":
		f.TipoPessoa.SetActive(0)
	case "Juridica":
		f.TipoPessoa.SetActive(1)
	}
	
	f.EstadoCivil.SetActiveID(c.EstadoCivil)
	
	switch(c.EstadoCivil) {
		case "solteiro":
			f.EstadoCivil.SetActive(0)
		case "casado":
			f.EstadoCivil.SetActive(1)
		case "separado":
			f.EstadoCivil.SetActive(2)
		case "divorciado": 
			f.EstadoCivil.SetActive(3)
		case "viúvo":
			f.EstadoCivil.SetActive(4)
	}

	f.Telefone.SetText(c.Telefone)
	f.WhatsApp.SetActive(c.Whatsapp)
	f.WhatsAppBtn.SetSensitive(c.Whatsapp)
	f.TelefoneAlt.SetText(c.TelefoneAlt)
	f.Email.SetText(c.Email)
	f.PaisNatal.SetText(c.PaisNatal)
	f.EstadoNatal.SetText(c.EstadoNatal)
	f.CidadeNatal.SetText(c.CidadeNatal)
	f.Project = make([]ProjectFields, len(c.Projects))

	//Interactions
	f.WhatsAppBtn.Connect("clicked", func ()  {
		url := handler.GetPreFormatedWhatsappUrl(getDataFromEntry(f.Telefone))
		handler.OpenInBrowser(url)
	})

	for i, p := range c.Projects {
		projectName := fmt.Sprint("Projeto localizado em >> " + p.Cidade + ", " + p.Bairro + ", ", p.Endereco + " Nº ", p.Numero)
		project, err := gtk.ExpanderNew(projectName)
		handler.Error("ui/Result.go >> client, gtk.BoxNew: ", err)

		form, fields := makeproject()
		
		fields.Cep.SetText(p.CEP)
		fields.Estado.SetText(p.Estado)
		fields.Cidade.SetText(p.Cidade)
		fields.Bairro.SetText(p.Bairro)
		fields.Endereco.SetText(p.Endereco)
		fields.Numero.SetText(fmt.Sprint(p.Numero))
		fields.Complemento.SetText(p.Complemento)
		switch (p.Status) {
		case "Inicial":
			fields.Status.SetActive(0)
		case "Pagamento Inicial Confirmado":
			fields.Status.SetActive(1)
		case "Em produção":
			fields.Status.SetActive(2)
		case "Instalado":
			fields.Status.SetActive(3)
		case "Pagamento Final Confirmado":
			fields.Status.SetActive(4)
		case "Finalizado":
			fields.Status.SetActive(5)
		}

		b, _ := fields.Observacoes.GetBuffer()
		b.SetText(p.Observacoes)
		
		//LoadData for tree view
		store, valorEntry, payCombo, obsEntry, addButton, vl := widgets.PreFormForPay(form, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION)

		for _, pay := range p.Payments {
			widgets.AddRow(store, COLUMN_PAYMENT_TYPE, COLUMN_VALUE, COLUMN_OBSERVATION,
			pay.Way, handler.ConvertFloatIntoString(pay.Value), pay.Observation, vl)
		}

		fields.Payment.PayCombo = payCombo
		fields.Payment.ValueEntry = valorEntry
		fields.Payment.ObsEntry = obsEntry
		fields.Payment.AddButton = addButton
		fields.Payment.Store = store

		fields.Contrato.SetActive(p.Contrato)

		fields.Cep.Connect("activate", func ()  {
			address, err := cep.GetText()
			handler.Error("controller/ClientController.go >> address, cep.GetText()", err)
			cepMap := handler.GetAddress(address)
			fields.Cidade.SetText(cepMap.Localidade)
			fields.Bairro.SetText(cepMap.Bairro)
			fields.Endereco.SetText(cepMap.Logradouro)
			fields.Estado.SetText(cepMap.UF)
			fields.Numero.SetText("")
			fields.Complemento.SetText("")
		})



		fields.Enviroments = make([]EnviromentFields,len(p.Enviroments))

		for i2, e := range p.Enviroments {
			ex, env := makeEnviroment()

			env.Nome.SetText(e.Name)
			buff, _ := env.Materiais.GetBuffer()
			buff.SetText(e.Materials)
			env.Fabricacao.SelectDay(uint(e.Production.Day()))
			env.Fabricacao.SelectMonth(uint(e.Production.Month()), uint(e.Production.Year()))
			env.Instalacao.SelectDay(uint(e.Installation.Day()))
			env.Instalacao.SelectMonth(uint(e.Installation.Month()), uint(e.Installation.Year()))
			form.PackStart(ex, true, false, 10)
			fields.Enviroments[i2] = env

			env.Nome.Connect("changed", func ()  {
				ex.SetLabel(getDataFromEntry(env.Nome))
			})
		}

		f.Project[i] = fields
		project.Add(form)
		handlers.PackEnd(project, true, false, 5)
	}

	editMode(false, f)

	edit.Connect("clicked", func ()  {
		if !buttonState && thisPage == notebook.GetCurrentPage() {
			editMode(true, f)
			buttonState = true
			image, err := gtk.ImageNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
			handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
			edit.SetImage(image)
		} else {
			model := getModelResult(f, c)
			model.Save(dbResult)
			editMode(false, f)
			buttonState = false
			image, err := gtk.ImageNewFromIconName("document-edit-symbolic", gtk.ICON_SIZE_BUTTON)
			handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
			edit.SetImage(image)
		}
	})

	notebook.Connect("switch-page", func (_ *gtk.Notebook, _ *gtk.Widget, index int)  {
		if index == thisPage {
			if buttonState {
				editMode(true, f)
				buttonState = true
				image, err := gtk.ImageNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
				handler.Error("controller/ResultController.go >> edit.Connect() >> image new from icon name", err)
				edit.SetImage(image)
			}
		} else {
			editMode(false, f)
			image, err := gtk.ImageNewFromIconName("document-edit-symbolic", gtk.ICON_SIZE_BUTTON)
			handler.Error("controller/ResultController.go >> notebook.Connect() >> image new from icon name", err)
			edit.SetImage(image)
		}
	})
}

func makeproject() (*gtk.Box, ProjectFields) {
	ProjectForm, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Result.go >> clientForm, gtk.BoxNew: ", err)
	cep	:= widgets.PreForm("CEP", ProjectForm)
	cidade := widgets.PreForm("Cidade", ProjectForm)
	estado := widgets.PreForm("Estado", ProjectForm)
	bairro := widgets.PreForm("Bairro", ProjectForm)
	endereco := widgets.PreForm("Endereço", ProjectForm)
	numero := widgets.PreForm("Número", ProjectForm)
	complemento := widgets.PreForm("Complemento", ProjectForm)
	status := widgets.PreFormComboBox("Status do projeto", 
	[]string {"Inicial", "Pagamento Inicial Confirmado", "Em produção", "Instalado", "Pagamento Final Confirmado", "Finalizado"}, 
	ProjectForm)
	observacoes := widgets.PreFormTextView("Observações", ProjectForm)
	valor := widgets.PreForm("Valor do projeto", ProjectForm)
	contrato := widgets.PreFormCheckBox("Projeto por contrato", ProjectForm)
	
	fields := ProjectFields {
		Cep: cep,
		Cidade:      cidade,
		Estado:      estado,
		Bairro:      bairro,
		Endereco:    endereco,
		Numero:      numero,
		Complemento: complemento,
		Status:      status,
		Observacoes: observacoes,
		Valor:       valor,
		Contrato:    contrato,
		Enviroments: []EnviromentFields{},
	}

	return ProjectForm, fields
}

func makeEnviroment() (*gtk.Expander, EnviromentFields) {
	expander, err := gtk.ExpanderNew("Ambiente")
	handler.Error("ui/Project.go >> addExpanderForEnviroment >> expander, gtk.Expander", err)
	
	form, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	handler.Error("ui/Project.go >> addExpanderForEnviroment >> form, gtk.Box", err)
	
	value := widgets.PreForm("Nome do Ambiente", form)
	materiais := widgets.PreFormTextView("Materiais", form)
	fabricacao := widgets.PreFormCalendar("Data de fabricação", form)
	instalacao := widgets.PreFormCalendar("Data de instalação", form)
	
	value.Connect("changed", func() {
		s, _ := value.GetText()
		expander.SetLabel(s)
	})

	env := EnviromentFields{value, materiais, fabricacao, instalacao}
	expander.Add(form)
	return expander, env
}

func getModelResult(c ClientFields, clientDB database.Client) database.Client {
	client := database.Client{}
	client.ID = clientDB.ID
	client.Nome = getDataFromEntry(c.Nome)
	client.Cpf = getDataFromEntry(c.CPF)
	client.Rg = getDataFromEntry(c.RG)
	client.Nascimento = getDataFromCalendar(c.Nascimento)
	client.Sexo = getDataFromComboBox(c.Sexo)
	client.TipoPessoa = getDataFromComboBox(c.TipoPessoa)
	client.EstadoCivil = getDataFromComboBox(c.EstadoCivil)
	client.Telefone = getDataFromEntry(c.Telefone)
	client.Whatsapp = c.WhatsApp.GetActive()
	client.TelefoneAlt = getDataFromEntry(c.TelefoneAlt)
	client.Email = getDataFromEntry(c.Email)
	client.PaisNatal = getDataFromEntry(c.PaisNatal)
	client.EstadoNatal = getDataFromEntry(c.EstadoNatal)
	client.CidadeNatal = getDataFromEntry(c.CidadeNatal)

	for i, p := range c.Project {
		np := database.Project{}
		np.ClientId = clientDB.ID
		np.ID = clientDB.Projects[i].ID
		np.Cidade = getDataFromEntry(p.Cidade)
		np.Estado = getDataFromEntry(p.Estado)
		np.Bairro = getDataFromEntry(p.Bairro)
		np.Endereco = getDataFromEntry(p.Endereco)
		np.Numero = uint(toInt(getDataFromEntry(p.Numero)))
		np.Complemento = getDataFromEntry(p.Complemento)
		np.Status = getDataFromComboBox(p.Status)
		np.Observacoes = getDataFromTextView(p.Observacoes)
		np.ValorProjeto = toFloat(strings.ReplaceAll(getDataFromEntry(p.Valor), ",", "."))
		np.Contrato = p.Contrato.GetActive()

		for i2, e := range p.Enviroments {
			ne := database.Enviroment {
				ProjectID: clientDB.Projects[i].ID,
				Name: getDataFromEntry(e.Nome),
				Materials: getDataFromTextView(e.Materiais),
				Production: getDataFromCalendar(e.Fabricacao),
				Installation: getDataFromCalendar(e.Instalacao),
			}

			ne.ID = clientDB.Projects[i].Enviroments[i2].ID
			np.Enviroments = append(np.Enviroments, ne)
		}

		//Get data from the store
		p.Payment.Store.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			if !p.Payment.Store.IterIsValid(iter) {return false}
			paymentType, err := model.GetValue(iter, 0)
			if err != nil {return false}
			val, err := model.GetValue(iter, 1)
			if err != nil {return false}
			obs, err := model.GetValue(iter, 2)
			if err != nil {return false}


			way, err := paymentType.GetString()
			if err != nil {return false}
			value, err := val.GetString()
			if err != nil {return false}
			observation, err := obs.GetString()
			if err != nil {return false}

			npay := database.Payment {
				Value: handler.ConvertStringIntoFloat(value),
				Way: way,
				Observation: observation,
			}
			np.Payments = append(np.Payments, npay)
			return true
		})

		client.Projects = append(client.Projects, np)
	}
	return client
}

func editMode(isEditable bool, c ClientFields) {
	c.Nome.SetSensitive(isEditable)
	c.CPF.SetSensitive(isEditable)
	c.RG.SetSensitive(isEditable)
	c.Nascimento.SetSensitive(isEditable)
	c.Sexo.SetSensitive(isEditable)
	c.TipoPessoa.SetSensitive(isEditable)
	c.EstadoCivil.SetSensitive(isEditable)
	c.WhatsApp.SetSensitive(isEditable)
	c.TelefoneAlt.SetSensitive(isEditable)
	c.Telefone.SetSensitive(isEditable)
	c.Email.SetSensitive(isEditable)
	c.PaisNatal.SetSensitive(isEditable)
	c.EstadoNatal.SetSensitive(isEditable)
	c.CidadeNatal.SetSensitive(isEditable)

	for _, p := range c.Project {
		p.Cep.SetSensitive(isEditable)
		p.Cidade.SetSensitive(isEditable)
		p.Estado.SetSensitive(isEditable)
		p.Bairro.SetSensitive(isEditable)
		p.Endereco.SetSensitive(isEditable)
		p.Numero.SetSensitive(isEditable)
		p.Complemento.SetSensitive(isEditable)
		p.Status.SetSensitive(isEditable)
		p.Observacoes.SetSensitive(isEditable)
		p.Valor.SetSensitive(isEditable)
		p.Contrato.SetSensitive(isEditable)

		p.Payment.PayCombo.SetSensitive(isEditable)
		p.Payment.ValueEntry.SetSensitive(isEditable)
		p.Payment.ObsEntry.SetSensitive(isEditable)
		p.Payment.AddButton.SetSensitive(isEditable)

		for _, e := range p.Enviroments {
			e.Nome.SetSensitive(isEditable)
			e.Fabricacao.SetSensitive(isEditable)
			e.Instalacao.SetSensitive(isEditable)
			e.Materiais.SetSensitive(isEditable)
		}
	}
}