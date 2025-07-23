package usecase

import (
	"context"

	"github.com/alexey/firstApp/domain/models"
)

/* UserUseCases объединяет все сценарии работы с пользователем */
type UserUsecase interface {
	// GetUserAuthorities(ctx context.Context, input dto.UserDTO) (*UserAuthoritiesOutput, error)
	// OpenPathGuider(ctx context.Context, input dto.UserDTO) (*UserAuthoritiesOutput, error)

	GetUserAuthorities(ctx context.Context, input *models.User) (*UserAuthoritiesOutput, error)
	OpenPathGuider(ctx context.Context, input *models.User) (*UserAuthoritiesOutput, error)
}

/* GetUserAuthoritiesOutput - выходные данные авторизации */
type UserAuthoritiesOutput struct {
	Email      string `json:"email" validate:"required,email"`
	Role       string `json:"role"`
	Token      string `json:"token,omitempty"`
	ExpiresAt  int64  `json:"expires_at,omitempty"`
	IsSelected string `json:"isselected"`
}
