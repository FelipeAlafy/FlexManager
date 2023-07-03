package database

import (
	"fmt"
	"github.com/FelipeAlafy/Flex/handler"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

func Connect(USER, PASS, HOST, PORT, DB string) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DB)
	db, err := gorm.Open("mysql", url)
	handler.Error("database/gorm.go >> db", err)
	return db
}