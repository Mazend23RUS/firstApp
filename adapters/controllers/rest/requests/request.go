package requests

import (
	"github.com/alexey/boundary/dto"
	userDTO "github.com/alexey/boundary/dto"
)

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	IsSelected bool   `json:"is_selected"`
}

func (lr *LoginRequest) MapperOfRequestToDTO() *userDTO.UserDTO {
	return dto.NewUserDTO(lr.Email, lr.Password, lr.IsSelected)
}

func (pe *ParseError) Error() string {
	return "Неверный JSON формат" + pe.Err.Error()
}

func (v *ValidateError) Error() string {
	return "Не валидные значения JSON запроса" + v.Err.Error()
}

/* Кастомные типы ошибок */
type ParseError struct{ Err error }
type ValidateError struct{ Err error }
