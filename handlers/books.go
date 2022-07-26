package handlers

import (
	"fmt"
	"net/http"
	"rest_api/models"

	"github.com/labstack/echo/v4"
)

//para validar criação
type CreateBookInput struct {
	Titulo string `json:"titulo" binding:"required"` //binding -> campo obrigatorio
	Autor  string `json:"autor" binding:"required"`
}

func FindBooks(context echo.Context) error {
	var books []models.Book
	models.DB.Find(&books) //select * from

	return context.JSON(http.StatusOK, books)

}

func CreateBook(context echo.Context) error {
	var input CreateBookInput
	//				  valida o resquest
	if err := context.Bind(&input); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	book := models.Book{Titulo: input.Titulo, Autor: input.Autor}
	err := models.DB.Create(&book).Error
	if err != nil {
		fmt.Println(err)
	}

	return context.JSON(http.StatusCreated, book)
}

func FindBook(context echo.Context) error {
	var book models.Book
	// 										"parametrizado"
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, "Nada encontrado!")
		return err
	}

	return context.JSON(http.StatusOK, book)
}

//para validar edição
type UpdateBookInput struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func UpdateBook(context echo.Context) error {

	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, map[string]string{
			"Erro": "Nenhum livro encontrado",
		})
		return err
	}

	var input UpdateBookInput
	if err := context.Bind(&input); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"Erro": err.Error(),
		})

	}

	updateBook := models.Book{Titulo: input.Titulo, Autor: input.Autor}

	if err := models.DB.Model(&book).Updates(&updateBook).Error; err != nil {
		context.JSON(500, err.Error())
		return err
	}

	return context.JSON(http.StatusOK, book)

}

func DeleteBook(context echo.Context) error {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, "livro não encontrado!")
		return err
	}

	models.DB.Delete(&book) // delete from
	// tratar o erro

	return context.JSON(http.StatusOK, "livro deletado")

}
