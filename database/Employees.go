package database

type Employees struct {
	Id uint `gorm:"primaryKey;auto_increment" json:"id"`
	Nome string `json:"nome"`
	Telefone string `json:"telefone"`
	Senha string `json:"senha"`
	Level uint `json:"level"`
}