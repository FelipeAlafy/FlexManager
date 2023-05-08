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
func (e Enviroment) Save(db *gorm.DB) {
	db.Save(&e)
}

func (p Project) SearchForEnviroments(db *gorm.DB) []Enviroment {
	enviroments := []Enviroment{}
	db.Where("project_id = ?", p.ID).Find(&enviroments)
	return enviroments
}