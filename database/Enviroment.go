package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Enviroment struct {
	gorm.Model
	ProjectID uint			`db:"project_id"`
	Name string				`db:"name"`
	Materials string		`db:"materials"`
	Production time.Time	`db:"production"`
	Installation time.Time	`db:"installation"`
}

func (e Enviroment) New(db *gorm.DB, p Project) {
	db.Save(&e)
}

//Edit Enviroment
func (old_env Enviroment) Update(db *gorm.DB, new_env Enviroment) {
	db.Model(&old_env).Update(new_env)
}