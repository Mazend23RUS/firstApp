package requests

import (
	"github.com/alexey/boundary/dto"
	userDTO "github.com/alexey/boundary/dto"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ButtonActiveRequest struct {
	Role      string
	Tokken    string
	Email     string
	Is_select IsSelected
}

// func (br *ButtonActiveRequest) ParsingAndValidateButtonRequest(body io.ReadCloser) error {
// 	defer body.Close()

// 	if err := json.NewDecoder(body).Decode(br); err != nil {
// 		return &ParseError{Err: err}
// 	}

// 	valid := validator.New()
// 	if err := valid.Struct(br); err != nil {
// 		return &ValidateError{Err: err}
// 	}
// 	return nil

// }

// func (lr *LoginRequest) ParsingAndValidateRequest(body io.ReadCloser) error {
// 	defer body.Close()

// 	/* 1. Парсинг JSON */
// 	if err := json.NewDecoder(body).Decode(lr); err != nil {
// 		return &ParseError{Err: err}
// 	}

// 	return nil
// }

func (lr *LoginRequest) MapperOfRequestToDTO() *userDTO.UserDTO {
	return dto.NewUserDTO(lr.Email, lr.Password)
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

/* Кастомные типы ошибок */
type ParseError struct{ Err error }
type ValidateError struct{ Err error }
type IsSelected struct{ IsSelected bool }
