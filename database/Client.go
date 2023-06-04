//Client.go
package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)


type Client struct {
	gorm.Model
	Nome string 			`db:"nome" json:"nome"`
	Rg string 				`db:"rg" json:"rg"`
	Cpf string 				`db:"cpf" json:"cpf"`
	Sexo string 			`db:"sexo" json:"sexo"`
	TipoPessoa string 		`db:"tipo_pessoa" json:"tipo_pessoa"`
	EstadoCivil string 		`db:"estado_civil" json:"estado_civil"`
	Nascimento time.Time 	`db:"nascimento" json:"nascimento"`
	Telefone string 		`db:"telefone" json:"telefone"`
	TelefoneAlt string 		`db:"telefone_alt" json:"telefone_alt"`
	Whatsapp bool 			`db:"whatsapp" json:"whatsapp"`
	Email string 			`db:"email" json:"email"`
	PaisNatal string 		`db:"pais_natal" json:"pais_natal"`
	EstadoNatal string 		`db:"estado_natal" json:"estado_natal"`
	CidadeNatal string 		`db:"cidade_natal" json:"cidade_natal"`
	Projects []Project 		`db:"projects" json:"projects"`
}


func (c *Client) New(db *gorm.DB) {
	db.Create(&c)
}

func (c *Client) Save(db *gorm.DB) {
	db.Save(&c)
	for _, p := range c.Projects {
		p.Save(db)
	}
}

func GetAllClients(db *gorm.DB) []Client {
	clients := []Client{}
	db.Find(&clients)
	return clients
}

func (c Client) Search(db *gorm.DB) []Client {
	clients := []Client{}
	name := fmt.Sprint("%", c.Nome, "%")
	db.Where("nome LIKE ?", name).Find(&clients)

	for c := 0; c < len(clients); c++ {
		clients[c].Projects = clients[c].SearchProjects(db)
	}

	return clients
}

func Search(db *gorm.DB, id uint) Client {
	client := Client{}
	db.First(&client, id)
	client.Projects = client.SearchProjects(db)
	return client
}

func (c Client) AddProject(db *gorm.DB, p Project) {
	c.Projects = append(c.Projects, p)
	db.Save(&c)
}
