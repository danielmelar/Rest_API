package main

import (
	"rest_api/handlers"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin router
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/books", handlers.FindBooks)
	router.POST("/books", handlers.CreateBook)
	router.GET("/books/:id", handlers.FindBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	router.Run()
}
