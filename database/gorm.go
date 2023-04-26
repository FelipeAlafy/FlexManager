package database

import (
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)



const (
	USER="root"
	PASS="@Flex1020"
	HOST="127.0.0.1"
	PORT="3306"
	DB="flex"
)

func Connect() *gorm.DB {
	url := "root:Flex1020@tcp(127.0.0.1:3306)/flex?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	handler.Error("database/gorm.go >> db", err)
	return db
}