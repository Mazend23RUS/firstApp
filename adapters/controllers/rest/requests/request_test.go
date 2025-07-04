package requests

import (
	"testing"

	"github.com/alexey/boundary/dto"
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
