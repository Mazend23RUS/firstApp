package readerequests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	errors_domain "github.com/alexey/firstApp/boundary/domain/errors"
	"github.com/stretchr/testify/assert"
)

func TestResponseWriter_SucssesResponse(t *testing.T) {

	tests := []struct {
		name     string
		data     interface{}
		expected string
	}{
		{
			name:     "Тестируем удачный ответ",
			data:     struct{ Message string }{Message: "success"},
			expected: `{"Message":"success"}`,
		},
	}

	writen := NewResponseWriter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			writen.SuccessResponse(w, tt.data)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, tt.expected, w.Body.String())
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		})
	}

}

func TestResponseWriter_ErrorResponse(t *testing.T) {

	tests := []struct {
		name         string
		err          error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Тест неудачного ответа ",
			err:          errors_domain.ErrValidationFailed,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"validation failed"}`,
		},
		{
			name:         "Юзер не найден ",
			err:          errors_domain.ErrUserNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: `{"error":"user not found"}`,
		},
		{
			name:         "Ошибка сервака ",
			err:          errors.New("internal server error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"internal server error"}`,
		},
	}

	rw := NewResponseWriter()

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			rw.ErrorResponse(w, tt.err)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		})
	}

}
