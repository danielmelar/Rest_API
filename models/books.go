package models

type Book struct {
	ID     uint   `json:"id"` //id para diferenciar cada livro no bd
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
