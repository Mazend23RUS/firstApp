package readerequests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/alexey/firstApp/boundary/dto"
	"github.com/stretchr/testify/assert"
)

func TestJSONRequestReader_ReadLogginRequest(t *testing.T) {

	tests := []struct {
		name        string
		requestbody string
		wanterr     bool
		expected    *dto.UserDTO
	}{
		{
			name:        "Тест чтения запроса",
			requestbody: `{"email":"test@testemail.com","password":"passwd123","is_selected":true}`,
			wanterr:     false,
			expected: &dto.UserDTO{
				Email:      "test@testemail.com",
				Password:   "passwd123",
				IsSelected: true,
			},
		},

		{
			name:        "invalid json",
			requestbody: `{"email":"test@example.com","password":`,
			wanterr:     true,
		},
	}

	reader := NewJSONRequestReader()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(tt.requestbody))
			req.Header.Set("Content-Type", "application/json")

			result, err := reader.ReadLoginRequest(req)
			if tt.wanterr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})

	}

}
