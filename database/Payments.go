package database

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	ProjectID uint			`db:"project_id"`
	Value float64
	Way string
	Observation string
}

func (p Payment) Save(db *gorm.DB) {
	db.Save(&p)
}

func (p Project) SearchForPayments(db *gorm.DB) []Payment {
	payments := []Payment{}
	db.Where("project_id = ?", p.ID).Find(&payments)
	return payments
}