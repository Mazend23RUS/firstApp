package usecase

import (
	"context"

	"github.com/alexey/boundary/dto"
)

// UserUseCases объединяет все сценарии работы с пользователем
type UserUsecase interface {
	GetUserAuthorities(ctx context.Context, input dto.UserDTO) (*UserAuthoritiesOutput, error)
	OpenPathGuider(ctx context.Context, isOpenGuider bool) error
}

// // GetUserAuthoritiesUseCase - сценарий получения авторизационных данных пользователя
// type GetUserAuthoritiesUseCase interface {
//     Execute(ctx context.Context, input GetUserAuthoritiesInput) (*GetUserAuthoritiesOutput, error)
// }

// // ComandToOpenPathguiderUseCase - сценарий открытия проводника
// type ComandToOpenPathguiderUseCase interface {
//     Execute(ctx context.Context, isOpenGuider bool) error
// }

// GetUserAuthoritiesInput - входные данные для авторизации
// type UserAuthoritiesInput struct {
// Email    string `json:"email" validate:"required,email"`
// Password string `json:"password" validate:"required,min=8"`

// }

// GetUserAuthoritiesOutput - выходные данные авторизации
type UserAuthoritiesOutput struct {
	Email     string `json:"email"`
	Role      string `json:"role"`
	Token     string `json:"token,omitempty"`
	ExpiresAt int64  `json:"expires_at,omitempty"`
}
