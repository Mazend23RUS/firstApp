package requests

import (
	"testing"

	"github.com/alexey/boundary/dto"
	"github.com/alexey/infrastructure/http/validator"
	"github.com/stretchr/testify/assert"
)

func TestRequestFor_MappingRequestToDTO(t *testing.T) {
	tests := []struct {
		name     string
		requerst LoginRequest
		expected *dto.UserDTO
	}{
		{
			name: "Тест маппинга в ДТО",
			requerst: LoginRequest{
				Email:      "test@first.com",
				Password:   "passwrd123",
				IsSelected: true,
			},
			expected: &dto.UserDTO{
				Email:      "test@first.com",
				Password:   "passwrd123",
				IsSelected: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.requerst.MapperOfRequestToDTO()
			assert.Equal(t, tt.expected, result)

		})

	}

}

func TestRequestFor_MappingRequestToDTO_Negative(t *testing.T) {

	tests := []struct {
		name          string
		request       LoginRequest
		expected      *dto.UserDTO
		expectedError bool
	}{
		{
			name: "Пустой email ",
			request: LoginRequest{
				Email:      "",
				Password:   "passwrd123",
				IsSelected: true,
			},
			expectedError: true,
		},
		{
			name: "Неверный адрес",
			request: LoginRequest{
				Email:      "bboy23test.com",
				Password:   "passwrd123",
				IsSelected: true,
			},
			expectedError: true,
		},
		{
			name: "Короткий пароль ",
			request: LoginRequest{
				Email:      "test@invalid.com",
				Password:   "pas",
				IsSelected: true,
			},
			expectedError: true,
		},
		{
			name: "Пустой пароль ",
			request: LoginRequest{
				Email:      "test@invalid.com",
				Password:   "",
				IsSelected: true,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Здесь мы тестируем только маппинг, валидация должна быть отдельно
			dto := tt.request.MapperOfRequestToDTO()

			// Проверяем что маппинг выполнился, даже с невалидными данными
			assert.NotNil(t, dto)
			assert.Equal(t, tt.request.Email, dto.Email)
			assert.Equal(t, tt.request.Password, dto.Password)
			assert.Equal(t, tt.request.IsSelected, dto.IsSelected)

			// Дополнительная проверка - тестируем что валидация DTO fails
			validator := validator.NewValidator()
			err := validator.Validate(dto)
			if tt.expectedError {
				assert.Error(t, err, "Ожидалась ошибка валидации для "+tt.name)
			} else {
				assert.NoError(t, err, "Не ожидалась ошибка валидации для "+tt.name)
			}
		})
	}

}
