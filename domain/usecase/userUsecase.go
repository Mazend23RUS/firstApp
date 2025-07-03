package implementationUseCase

import (
	"context"
	"fmt"
	"time"

	usecase "github.com/alexey/boundary/domain/useCase"
	"github.com/alexey/boundary/dto"
	"github.com/alexey/pkg/logger"
)

type AuthUseCase struct {
	log logger.Logger
}

func NewAuthUseCase(log logger.Logger) *AuthUseCase {
	return &AuthUseCase{
		log: log,
	}
}

func (us *AuthUseCase) GetUserAuthorities(c context.Context, input dto.UserDTO) (*usecase.UserAuthoritiesOutput, error) {

	if input.Email == "bboy23@mail.ru" && input.Password == "87654321" {
		us.log.PrintInfo(c, "Введены данные Admin"+" email: "+input.Email)
		return &usecase.UserAuthoritiesOutput{
			Email:     input.Email,
			Role:      "Admin",
			Token:     "tokke-15654-5631-45$",
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		}, nil
	}

	us.log.PrintInfo(c, "Введены данные для user")
	return &usecase.UserAuthoritiesOutput{
		Email:     input.Email,
		Role:      "user",
		Token:     "168456ewq",
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
	}, nil
}

func (us *AuthUseCase) OpenPathGuider(c context.Context, input dto.UserDTO) (*usecase.UserAuthoritiesOutput, error) {
	if input.IsSelected == true {
		us.log.PrintInfo(c, "Нажата кнопка открытия проводника")

		return &usecase.UserAuthoritiesOutput{
			Email:      input.Email,
			Token:      "tokke-15654-5631-46$",
			ExpiresAt:  time.Now().Add(12 * time.Hour).Unix(),
			IsSelected: "Поисковик открылся",
		}, nil
	}
	return nil, fmt.Errorf("Ошибка при открытии окна %v", c)
}
