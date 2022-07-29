package main

import (
	"rest_api/handlers"
	"rest_api/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// server echo
	server := echo.New()

	// para gerar logs e controle de panics
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	models.Connection() // conexão com o banco

	server.GET("/books", handlers.FindBooks)
	server.POST("/books", handlers.CreateBook)
	server.GET("/books/:id", handlers.FindBook)
	server.PUT("/books/:id", handlers.UpdateBook)
	server.DELETE("/books/:id", handlers.DeleteBook)

	server.Logger.Fatal(server.Start(":8080"))
}
