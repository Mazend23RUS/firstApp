package implementationUseCase

import (
	"context"
	"time"

	usecase "github.com/alexey/boundary/domain/useCase"
	"github.com/alexey/boundary/dto"
	"github.com/alexey/pkg/logger"
)

type AuthUseCase struct {
	log logger.Logger
	// AuthUseCase usecase.UserUsecase
}

func NewAuthUseCase(log logger.Logger) *AuthUseCase {
	return &AuthUseCase{
		log: log,
	}
}

func (us *AuthUseCase) GetUserAuthorities(c context.Context, input dto.UserDTO) (*usecase.UserAuthoritiesOutput, error) {

	// if input.IsSelected == true {
	// 	if err := us.OpenPathGuider(c, input); err != nil {
	// 		return &usecase.UserAuthoritiesOutput{
	// 			Email:      input.Email,
	// 			Token:      "tokke-15654-5631-45$",
	// 			ExpiresAt:  time.Now().Add(12 * time.Hour).Unix(),
	// 			IsSelected: "Поисковик открылся",
	// 		}, fmt.Errorf("Ошибка открытия окна", err)
	// 	}
	// }

	if input.Email == "bboy23@mail.ru" && input.Password == "87654321" {
		us.log.PtintInfo(c, "Введены данные Admin"+" email: "+input.Email)
		return &usecase.UserAuthoritiesOutput{
			Email:     input.Email,
			Role:      "Admin",
			Token:     "tokke-15654-5631-45$",
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		}, nil
	}

	us.log.PtintInfo(c, "Введены данные для user")
	return &usecase.UserAuthoritiesOutput{
		Email:     input.Email,
		Role:      "user",
		Token:     "168456ewq",
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
	}, nil
}

func (us *AuthUseCase) OpenPathGuider(c context.Context, input dto.UserDTO) *usecase.UserAuthoritiesOutput {
	if input.IsSelected == true {
		us.log.PtintInfo(c, "Нажата кнопка открытия проводника")
		return &usecase.UserAuthoritiesOutput{
			Email:      input.Email,
			Token:      "tokke-15654-5631-46$",
			ExpiresAt:  time.Now().Add(12 * time.Hour).Unix(),
			IsSelected: "Поисковик открылся",
		}
	}
	return nil
}
