package handlers

import (
	"net/http"
	"rest_api/models"

	"github.com/labstack/echo/v4"
)

func FindBooks(context echo.Context) error {
	var books []models.Book
	if err := models.DB.Find(&books).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, books)
	return nil

}

// para validar os dados
type CreateBookInput struct {
	Titulo string `json:"titulo" binding:"required"` //binding -> campo obrigatorio
	Autor  string `json:"autor" binding:"required"`
}

func CreateBook(context echo.Context) error {
	var input CreateBookInput

	if err := context.Bind(&input); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	book := models.Book{Titulo: input.Titulo, Autor: input.Autor}
	err := models.DB.Create(&book).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	context.JSON(http.StatusCreated, echo.Map{
		"Livro criado": book,
	})
	return nil
}

func FindBook(context echo.Context) error {
	var book models.Book

	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		return context.JSON(http.StatusBadRequest, "Nada encontrado!")

	}

	context.JSON(http.StatusOK, book)
	return nil
}

//para validar edição
type UpdateBookInput struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func UpdateBook(context echo.Context) error {

	var book models.Book

	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		return context.JSON(http.StatusBadRequest, "Livro não encontrado!")

	}

	var input UpdateBookInput

	if err := context.Bind(&input); err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"Erro": err.Error(),
		})

	}

	updateBook := models.Book{Titulo: input.Titulo, Autor: input.Autor}

	if err := models.DB.Model(&book).Where("id = ?", context.Param("id")).Updates(&updateBook).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())

	}

	context.JSON(http.StatusOK, book)
	return nil

}

func DeleteBook(context echo.Context) error {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, "livro não encontrado!")
		return err
	}

	// delete from
	if err := models.DB.Delete(&book); err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, "livro deletado")
	return nil

}
