package database

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	ClientId		uint
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
}

func GetAllProjects(db *gorm.DB) []Project {
	projects := []Project{}
	db.Find(&projects)
	return projects
}

func (c Client) SearchProjects(db *gorm.DB) Client {
	projects := []Project{}
	db.Where("client_id = ?", c.ID).Find(&projects)
	c.Projects = projects
	return c
}