package model

type Expense struct {
	Id              int `gorm:"primaryKey;autoIncrement"`
	UserId          int `gorm:"column:userId"`
	Description     string
	Amount          float64
	SpendFrom       string
	Date            string `gorm:"column:spend_date"`
	Category        string
	AdditionalNotes string
	Image           string
	Tag             string
	IsCounted       bool
	AddedAs         string
	IsReviewed      bool
}

