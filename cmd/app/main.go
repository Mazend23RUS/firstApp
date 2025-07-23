package main

import (
	"context"

	"github.com/alexey/firstApp/adapters/controllers"
	renderequests "github.com/alexey/firstApp/adapters/controllers/rest"

	infrahttp "github.com/alexey/firstApp/infrastructure/http" // Общий алиас

	implementationUseCase "github.com/alexey/firstApp/domain/usecase"
	"github.com/alexey/firstApp/pkg/logger"
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

	/* NewController() инициализация контроллера */
	contro := controllers.NewController(log, authUseCase, readr, respo)

	/* NewGinServer() инициализация сервера */
	ser := infrahttp.NewGinServer(log)

	/* SetupRouter() Инициализация router */
	infrahttp.SetupRouter(ser, contro)

	log.PrintInfo(context.Background(), "Попытка стартануть сервер ")
	if err := ser.Start(port); err != nil {
		log.PrintError(context.TODO(), "Сервер не стартанул", err)
	}

}
