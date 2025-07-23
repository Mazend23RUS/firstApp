package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexey/firstApp/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestServer_RegisterPublicRoute(t *testing.T) {
	server := NewGinServer(logger.InitLogger())

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server.RegisterPublicRoute("GET", "/", handler)

	req := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	server.engin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestServer_GinStartServer(t *testing.T) {

	server := NewGinServer(logger.InitLogger())

	// Регистрируем тестовый маршрут
	server.RegisterPublicRoute("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Создаем тестовый HTTP сервер
	ts := httptest.NewServer(server.engin)
	defer ts.Close()

	// Проверяем работу сервера
	resp, err := http.Get(ts.URL + "/test")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "OK", string(body))
}
