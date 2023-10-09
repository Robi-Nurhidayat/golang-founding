package main

import (
	"bwa/golang/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "root:@tcp(localhost:3306)/bwastarup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = " test simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "band"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
