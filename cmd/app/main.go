package main

import (
	"context"
	"fmt"

	"github.com/alexey/firstApp/adapters/controllers"

	"github.com/alexey/firstApp/adapters/controllers/fanout"
	renderequests "github.com/alexey/firstApp/adapters/controllers/rest"

	infrahttp "github.com/alexey/firstApp/infrastructure/http" // Общий алиас

	implementationUseCase "github.com/alexey/firstApp/domain/usecase"
	"github.com/alexey/firstApp/pkg/logger"
)

func main() {

	port := ":8080"

	/* InitLogger() Инициализация логгера */
	log := logger.InitLogger()

	/* NewZapLogger() Инициализация zap логгера */
	zaplog, _ := logger.NewZapLogger()

	// Регистрация логгеров
	logReg := logger.NewRegisteredLog()
	logReg.RegisterLogger("stdlog", log)
	logReg.RegisterLogger("zap", zaplog)

	fmt.Println("смотрим что лежит в мапе ", logReg.GetAllRegLogger())

	messageChan := make(chan interface{}, 100)
	splitter := fanout.NewSplitter[interface{}](100, logReg)

	// Запуск обработки входящих сообщений
	splitter.Start(messageChan)

	/* NewJSONRequestReader() инициализация читателя запросов */
	readr := renderequests.NewJSONRequestReader()

	/* NewResponseWriter() инициализация ответа */
	respo := renderequests.NewResponseWriter()

	/* NewAuthUseCase() Инициализация use case */
	authUseCase := implementationUseCase.NewAuthUseCase(log)

	/* NewController() инициализация контроллера */
	contro := controllers.NewController(log, authUseCase, readr, respo, messageChan)

	/* NewGinServer() инициализация сервера */
	ser := infrahttp.NewGinServer(log)

	/* SetupRouter() Инициализация router */
	infrahttp.SetupRouter(ser, contro)

	log.PrintInfo(context.Background(), "Попытка стартануть сервер ")
	if err := ser.Start(port); err != nil {
		log.PrintError(context.TODO(), "Сервер не стартанул", err)
	}

}
