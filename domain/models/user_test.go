package models_test

import (
	"testing"

	"github.com/alexey/firstApp/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		fullName    string
		wantErr     bool
		errContains string
	}{
		{
			name:     "валидный user",
			email:    "test@example.com",
			password: "SecurePass123",
			fullName: "Alexey Petrov",
			wantErr:  false,
		},
		{
			name:        "невалидный email",
			email:       "invalid email",
			password:    "SecurePass123",
			fullName:    "Alexey Petrov",
			wantErr:     true,
			errContains: "not valid email",
		},
		{
			name:        "короткий пароль",
			email:       "test@email.com",
			password:    "short",
			fullName:    "Alexey Petrov",
			wantErr:     true,
			errContains: "password validation",
		},

		{
			name:        "отсутствие цифры",
			email:       "test@email.com",
			password:    "shortpassword",
			fullName:    "Alexey Petrov",
			wantErr:     true,
			errContains: "Number",
		},

		{
			name:        "Имя без заглавной буквы",
			email:       "test@email.com",
			password:    "SecurePass123",
			fullName:    "alexey",
			wantErr:     true,
			errContains: "first symbol must be Upper",
		},
	}

	for _, tes := range tests {
		t.Run(tes.name, func(t *testing.T) {
			user, err := models.NewUser(tes.email, tes.password, tes.fullName, nil, false)

			if tes.wantErr {
				assert.Error(t, err)
				if tes.errContains != "" {
					assert.Contains(t, err.Error(), tes.errContains)
				}
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tes.email, string(user.Email()))
			assert.Equal(t, tes.fullName, string(user.Name()))
		})
	}
}
