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

	return db
}