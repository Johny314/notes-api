package models

type Note struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"required"`
}
