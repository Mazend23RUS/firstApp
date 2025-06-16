package implementationUseCase

import (
	"context"
	"fmt"
	"time"

	usecase "github.com/alexey/boundary/domain/useCase"
	"github.com/alexey/boundary/dto"
	"github.com/alexey/pkg/logger"
)

func (us *AuthUseCase) GetUserAuthorities(c context.Context, input dto.UserDTO) (*usecase.UserAuthoritiesOutput, error) {

	if input.Email == "bboy23@mail.ru" && input.Password == "87654321" {
		us.log.PtintInfo(c, "Введены данные Admin"+" email: "+input.Email+" password_Длинна: "+fmt.Sprint((len(input.Password))))
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

	return nil, fmt.Errorf("Неверные данные на входе")
}

func (uc *AuthUseCase) OpenPathGuider(c context.Context, isOpenGuider bool) error {

	if !isOpenGuider {
		return fmt.Errorf("guider not opened")
	}
	uc.log.PtintInfo(c, "Path guider opened successfully")
	return nil
}

func NewAuthUseCase(log logger.Logger) *AuthUseCase {
	return &AuthUseCase{
		log: log,
	}
}

type AuthUseCase struct {
	log      logger.Logger
	Email    string
	Password string
}
