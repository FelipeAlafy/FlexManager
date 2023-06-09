package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	ExpenseValue float64
	ExpenseType string
	ExpenseObservation string
}

func (e Expense) Save(db *gorm.DB) {
	db.Save(&e)
}

func GetPaymentsDataByDate(db *gorm.DB, startDate, endDate time.Time) []Payment {
	payments := []Payment{}
	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&payments)
	return payments
}

func GetExpensesDataByDate(db *gorm.DB, startDate, endDate time.Time) []Expense {
	expenses := []Expense{}
	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&expenses)
	return expenses
}

func GetCountOfProjects(db *gorm.DB, startDate, endDate time.Time) int {
	projects := []Project{}
	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&projects)
	return len(projects)
}

func GetCountOfProjectsByStatus(db *gorm.DB, startDate, endDate time.Time, status Project) int {
	projects := []Project{}
	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Where("status = ?", status.Status).Find(&projects)
	return len(projects)
}

func DeleteAllExpenses(db *gorm.DB, startDate, endDate *time.Time) {
	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Delete(&Expense{})
}