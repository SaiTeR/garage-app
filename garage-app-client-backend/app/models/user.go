package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone        string `gorm:"type:varchar(15);not null;unique"`
	Name         string `gorm:"type:varchar(100);not null"`
	Surname      string `gorm:"type:varchar(100);not null"`
	Password     string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(100);unique"`
	BonusBalance int    `gorm:"type:int;default:0"`
}

func (User) TableName() string {
	return "clients" // Переименовать таблицу
}
