package database

import (
	"github.com/jinzhu/gorm"
)

func Run() *gorm.DB {
	db := Connect()
	//Auto create and update tables
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Employees{})
	db.AutoMigrate(&Enviroment{})
	db.AutoMigrate(&Payment{})
	db.AutoMigrate(&Expense{})

	return db
}