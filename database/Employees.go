package database

import (
	"github.com/jinzhu/gorm"
)

type Employees struct {
	Id uint `gorm:"primaryKey;auto_increment" json:"id"`
	Nome string `json:"nome"`
	Telefone string `json:"telefone"`
	Senha string `json:"senha"`
	Level uint `json:"level"`
}

func GetAllEmployees(db *gorm.DB) []Employees {
	employees := []Employees{}
	db.Find(&employees)
	return employees
}

func CheckLogin(username, password string, db *gorm.DB) Employees {
	employee := []Employees{}
	db.Where(&Employees{Nome: username, Senha: password}).Find(&employee)
	return employee[0]
}

func (employee Employees) Save(db *gorm.DB) {
	db.Save(&employee)
}