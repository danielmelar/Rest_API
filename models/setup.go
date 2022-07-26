package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=1234"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := database.AutoMigrate(&Book{}); err != nil {
		fmt.Println(err)
	}

	DB = database

	fmt.Println("Conectado com sucesso!")

}
