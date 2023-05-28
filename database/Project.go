package database

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	ClientId		uint
	CEP				string
	Cidade 			string
	Estado 			string
	Bairro 			string
	Endereco 		string
	Numero 			uint
	Complemento 	string
	Status 			string
	Observacoes 	string
	ValorProjeto 	float64
	Contrato 		bool
	Enviroments		[]Enviroment
	Payments		[]Payment
}

func GetAllProjects(db *gorm.DB) []Project {
	projects := []Project{}
	db.Find(&projects)
	return projects
}

func (c Client) SearchProjects(db *gorm.DB) []Project {
	projects := []Project{}
	db.Where("client_id = ?", c.ID).Find(&projects)

	for p := 0; p < len(projects); p++ {
		projects[p].Enviroments = projects[p].SearchForEnviroments(db)
		projects[p].Payments = projects[p].SearchForPayments(db)
	}

	return projects
}

func (p Project) Save(db *gorm.DB) {
	db.Save(&p)
	for _, e := range p.Enviroments {
		e.Save(db)
	}
	for _, pm := range p.Payments {
		println(pm.Value)
		pm.Save(db)
	}
}