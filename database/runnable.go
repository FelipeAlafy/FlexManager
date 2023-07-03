package database

import (
	"github.com/jinzhu/gorm"
)

func Run(USER, PASS, HOST, PORT, DB string) *gorm.DB {
	db := Connect(USER, PASS, HOST, PORT, DB)
	//Auto create and update tables
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Employees{})
	db.AutoMigrate(&Enviroment{})
	db.AutoMigrate(&Payment{})
	db.AutoMigrate(&Expense{})

	return db
}