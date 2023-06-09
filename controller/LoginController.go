package controller

import (
	"github.com/FelipeAlafy/Flex/database"
	"github.com/jinzhu/gorm"
)

func Login(username, password string, dbUsers *gorm.DB) (bool, database.Employees) {
	employees := GetUsers(dbUsers)
	for _, e := range employees {
		if e.Nome == username && e.Senha == password {
			return true, e
		}
	}
	return false, database.Employees{}
}