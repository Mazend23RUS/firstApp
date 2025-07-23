package implementationUseCase

import (
	"context"
	"testing"

	usecase "github.com/alexey/firstApp/boundary/domain/useCase"
	"github.com/alexey/firstApp/domain/models"
	"github.com/alexey/firstApp/pkg/logger"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseImplementation_GetAuthorities(t *testing.T) {

	tests := []struct {
		name         string
		model        *models.User
		wanterr      bool
		expextedBody *usecase.UserAuthoritiesOutput
	}{
		{
			name: "Тестирование получения авторизации Admin",
			model: &models.User{
				Email:    "bboy23@mail.ru",
				Password: "87654321",
			},
			wanterr: false,
			expextedBody: &usecase.UserAuthoritiesOutput{
				Email: "bboy23@mail.ru",
				Token: "tokke-15654-5631-45$",
				Role:  "Admin",
			},
		},
		{
			name: "Тестирование получения авторизации User",
			model: &models.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			expextedBody: &usecase.UserAuthoritiesOutput{
				Email: "test@example.com",
				Token: "168456ewq",
				Role:  "user",
			},
		},
	}

	log := logger.InitLogger()
	authorities := NewAuthUseCase(log)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := authorities.GetUserAuthorities(context.Background(), tt.model)
			if tt.wanterr {
				assert.Error(t, err)

			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expextedBody.Email, result.Email)
				assert.Equal(t, tt.expextedBody.Role, result.Role)
				assert.NotEmpty(t, result.Token)
				assert.NotZero(t, result.ExpiresAt)

			}

		})

	}

}
