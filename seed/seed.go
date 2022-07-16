package main

import (
	"example/go-orm-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "warawit:admin@tcp(127.0.0.1:3306)/app_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "Ivy", Lname: "Cal", Username: "ivy.cal@mecallapi.com", Avatar: "https://www.mecallapi.com/users/2.png"})
	db.Create(&model.User{Fname: "Walter", Lname: "Beau", Username: "walter.beau@mecallapi.com", Avatar: "https://www.mecallapi.com/users/3.png"})
}
