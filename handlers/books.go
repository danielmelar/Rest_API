package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)

//para validar criação
type CreateBookInput struct {
	Titulo string `json:"titulo" binding:"required"` //binding -> campo obrigatorio
	Autor  string `json:"autor" binding:"required"`
}

//para validar edição
type UpdateBookInput struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func FindBooks(context *gin.Context) {
	var books []models.Book
	models.DB.Find(&books) //select * from

	context.JSON(http.StatusOK, gin.H{
		"livros": books})

}

func Verifica() error {
	return errors.New("Falha ao executar o método")
}

func CreateBook(context *gin.Context) {
	var input CreateBookInput
	//				  valida o resquest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	book := models.Book{Titulo: input.Titulo, Autor: input.Autor}
	err := models.DB.Create(&book).Error
	if err != nil {
		fmt.Println(err)
	}

	context.JSON(http.StatusCreated, gin.H{"livros": book})
}

func FindBook(context *gin.Context) {
	var book models.Book
	// 										"parametrizado"
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"erro": "Nada encontrado!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"livro:": book,
	})
}

func UpdateBook(context *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Nenhum livro encontrado!",
		})
		return
	}

	var input UpdateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
		return
	}

	models.DB.Model(&book).Updates(input) //update set
	// tratar o erro

	context.JSON(http.StatusOK, gin.H{
		"livros": book,
	})

}

func DeleteBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Erro":     err.Error(),
			"Mensagem": "livro não encontrado!",
		})
		return
	}

	models.DB.Delete(&book) // delete from
	// tratar o erro

	context.JSON(http.StatusOK, gin.H{
		"mensagem:": "livro deletado",
	})

}
