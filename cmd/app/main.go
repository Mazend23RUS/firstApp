package main

import (
	"context"

	"github.com/alexey/adapters/controllers"
	renderequests "github.com/alexey/adapters/controllers/rest"

	routing "github.com/alexey/infrastructure/http"
	server "github.com/alexey/infrastructure/http"
	"github.com/alexey/infrastructure/http/validator"

	implementationUseCase "github.com/alexey/domain/usecase"
	"github.com/alexey/pkg/logger"
)

func main() {

	port := ":8080"

	/* InitLogger() Инициализация логгера */
	log := logger.InitLogger()

	/* NewJSONRequestReader() инициализация читателя запросов */
	readr := renderequests.NewJSONRequestReader()

	/* NewResponseWriter() инициализация ответа */
	respo := renderequests.NewResponseWriter()

	/* NewAuthUseCase() Инициализация use case */
	authUseCase := implementationUseCase.NewAuthUseCase(log)

	/* NewValidator() Инициализация Валидатора */
	validat := validator.NewValidator()

	/* Инициализация errorHandler */
	errhand := renderequests.NewErrorStatus()

	/* NewController() инициализация контроллера */
	contro := controllers.NewController(log, validat, authUseCase, readr, respo, errhand)

	/* NewGinServer() инициализация сервера */
	ser := server.NewGinServer()

	/* SetupRouter() Инициализация router */
	routing.SetupRouter(ser, contro)

	log.PtintInfo(context.Background(), "Стартуем сервер на порте: 8080")
	if err := ser.Start(port); err != nil {
		log.PrintError(context.TODO(), "Сервер не стартанул", err)
	}

}
