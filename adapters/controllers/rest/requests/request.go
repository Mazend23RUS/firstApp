package requests

import (
	"encoding/json"
	"io"

	userDTO "github.com/alexey/boundary/dto"
	"github.com/go-playground/validator/v10"
)

// type CreateRequest struct {
// 	Name        string `json:"name" validate:"required,min=2,max=50"`
// 	Email       string `json:"email" validate:"required,email"`
// 	Password    string `json:"password" validate:"required,min=6"`
// 	PhoneNumber int    `json:"phone" validate:"required,min=12,max=12"`
// }

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (lr *LoginRequest) ParsingAndValidateRequest(body io.ReadCloser) error {
	defer body.Close()

	// 1. Парсинг JSON
	if err := json.NewDecoder(body).Decode(lr); err != nil {
		return &ParseError{Err: err}
	}
	// 2. Валидация структуры
	valid := validator.New()
	if err := valid.Struct(lr); err != nil {
		return &ValidateError{Err: err}
	}
	return nil
}

func (r *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (lr *LoginRequest) MapperOfRequestToDTO() userDTO.UserDTO {
	return userDTO.UserDTO{
		Email:    lr.Email,
		Password: lr.Password,
	}
}

func (pe *ParseError) Error() string {
	return "Неверный JSON формат" + pe.Err.Error()
}

func (v *ValidateError) Error() string {
	return "Не валидные значения JSON запроса" + v.Err.Error()
}

func (s *IsSelected) Selected() bool {
	return true
}

// Кастомные типы ошибок
type ParseError struct{ Err error }
type ValidateError struct{ Err error }
type IsSelected struct{ IsSelected bool }

// type MapperOfRequestToDTO struct {
//    CreateRequest

// }

// func MapperOfRequestToDTO(c *CreateRequest) userDTO.UserDTO {
//     user := userDTO.UserDTO{
// 		PhoneNumber : c.PhoneNumber,
// 		Email: c.Email,
//         Name: c.Name,
// 		Password: c.Password,
// 	}

//    return user
// }
