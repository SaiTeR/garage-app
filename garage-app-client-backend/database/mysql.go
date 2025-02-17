package database

import (
	"fmt"
	"garage-app-client-backend/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitMySQL() {
	dsn := "root:root@tcp(mysql:3306)/garage_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to MySQL")
	}
	fmt.Println("Successfully connected to MySQL")
}

func MigrateUsers() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
		panic("Migrate USERS: FAILED")
	}
	fmt.Println("Migrate USERS: success")
}

func GetMySQLClient() *gorm.DB {
	return DB
}
