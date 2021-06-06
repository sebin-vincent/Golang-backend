package model

type Category struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	TYPE      string `gorm:"column:type"`
	IsEnabled bool   `gorm:"column:isEnabled"`
}
