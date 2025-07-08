package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	readerequests "github.com/alexey/adapters/controllers/rest"
	implementationUseCase "github.com/alexey/domain/usecase"

	"github.com/alexey/pkg/logger"
)

func TestController_LoginHandler(t *testing.T) {
	// Инициализация зависимостей
	log := logger.InitLogger()
	reader := readerequests.NewJSONRequestReader()
	responder := readerequests.NewResponseWriter()
	authUseCase := implementationUseCase.NewAuthUseCase(log)
	controller := NewController(log, authUseCase, reader, responder)

	tests := []struct {
		name         string
		requestsbody string
		expecstatus  int
		isErrorCase  bool
	}{
		{
			name:         "Тестирование_контроллера",
			requestsbody: `{"email":"bboy23@mail.ru","password":"87654321","is_selected":false}`,
			expecstatus:  http.StatusOK,
			isErrorCase:  false,
		},
		{
			name:         "Плохой_запрос",
			requestsbody: `{"email":"bboy23@mail.ru","password":`,
			expecstatus:  http.StatusBadRequest,
			isErrorCase:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(tt.requestsbody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			controller.LoginHandler(w, req)

			if tt.isErrorCase {

				if w.Code != http.StatusBadRequest {
					t.Errorf("Для теста '%s' ожидался статус %d, получен %d. Тело ответа: %s",
						tt.name, http.StatusBadRequest, w.Code, w.Body.String())
				}
			} else {
				if w.Code != tt.expecstatus {
					t.Errorf("Для теста '%s' ожидался статус %d, получен %d",
						tt.name, tt.expecstatus, w.Code)
				}
			}
		})
	}
}

func TestController_OpenPathGuader(t *testing.T) {
	log := logger.InitLogger()
	requestreader := readerequests.NewJSONRequestReader()
	response := readerequests.NewResponseWriter()
	auth := implementationUseCase.NewAuthUseCase(log)

	controller := NewController(log, auth, requestreader, response)

	tests := []struct {
		name         string
		request      string
		responsecode int
		wanterr      bool
	}{
		{
			name:         "Успешное_открытие_проводника",
			request:      `{"email":"bboy23@mail.ru","password":"87654321","is_selected":true}`,
			responsecode: http.StatusOK,
			wanterr:      false,
		},
		{
			name:         "Проводник_не_открыт_(is_selected_false)",
			request:      `{"email":"bboy23@mail.ru","password":"87654321","is_selected":false}`,
			responsecode: http.StatusInternalServerError,
			wanterr:      true,
		},
		{
			name:         "Невалидный_JSON",
			request:      `{"email":"bboy23@mail.ru","password":`,
			responsecode: http.StatusBadRequest,
			wanterr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/button", bytes.NewBufferString(tt.request))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			controller.ButtonHandler(w, req)

			if tt.wanterr {
				// Для ошибок проверяем, что статус не 2xx
				if w.Code >= 200 && w.Code < 300 {
					t.Errorf("Для теста '%s' ожидалась ошибка, получен статус %d. Тело ответа: %s",
						tt.name, w.Code, w.Body.String())
				}
			} else {
				if w.Code != tt.responsecode {
					t.Errorf("Для теста '%s' ожидался статус %d, получен %d. Тело ответа: %s",
						tt.name, tt.responsecode, w.Code, w.Body.String())
				}
			}
		})
	}
}
