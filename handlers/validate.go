package handlers

import (
	"github.com/go-playground/validator/v10"
)

// Общий валидатор для всего пакета
var validate = validator.New()
