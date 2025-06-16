package main

import (
	"context"
	"net/http"

	"github.com/alexey/adapters/controllers"
	routing "github.com/alexey/infrastructure/http"

	implementationUseCase "github.com/alexey/domain/usecase"
	"github.com/alexey/pkg/logger"
	"github.com/gin-gonic/gin"
)

/*GinServer - структура, которая инкапсулирует сервер Gin. Содержит поле engine - экземпляр Gin-движка. */
type GinServer struct {
	engin *gin.Engine // gin сслка на пакет import, вызываем у пакета структуру Engine с кучей различных методов
}

/*NewGinServer() - конструктор, который:
Создаёт новый экземпляр Gin с помощью gin.New()
(в отличие от gin.Default(),
который автоматически добавляет Logger и Recovery middleware) */

func NewGinServer() *GinServer {

	server := &GinServer{
		engin: gin.New(),
	}

	// Добавляем глобальные middleware
	/* Явно добавляет два стандартных middleware: */
	server.engin.Use(gin.Logger())   // Logger() - логирует все HTTP-запросы
	server.engin.Use(gin.Recovery()) // Recovery() - перехватывает паники и возвращает 500 ошибку вместо падения сервера

	return server

}

//  ------- Реализация интерфейса Server --------

/*
	RegisterPublicRoute() - регистрирует обработчик для конкретного HTTP-метода и пути:

method - HTTP-метод (GET, POST и т.д.)
path - URL-путь
handler - стандартный обработчик http.HandlerFunc
gin.WrapH() адаптирует стандартный http.HandlerFunc к формату Gin
*/
func (s *GinServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	s.engin.Handle(method, path, gin.WrapF(handler))
}

/* Стартует сервер  Start() - запускает сервер на указанном адресе (например, ":8080") */
func (s *GinServer) Start(address string) error {
	return s.engin.Run(address)
}

/* AuthMiddleware() - фабрика middleware, возвращающая функцию-обработчик типа gin.HandlerFunc */

func AuthMidlleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		if g.GetHeader("Authorization") == "" { // Проверяется наличие заголовка Authorization
			g.AbortWithStatus(http.StatusUnauthorized) // Если заголовка нет - возвращается статус 401 (Unauthorized) и обработка прерывается (AbortWithStatus)
			return
		}
		g.Next() // Если заголовок есть - вызывается c.Next() для передачи управления следующему обработчику в цепочке
	}
	/* Middleware в Gin - это функции, которые получают контекст запроса и могут:
	   Модифицировать запрос
	   Проверять условия
	   Прерывать выполнение
	   Передавать управление следующему обработчику */
}

/* Инициализация usecase */
// NewAuthUseCase создает реализацию UserUseCases

// AuthUseCaseImpl реализует UserUseCases

// func initAuthUseCase(log logger.Logger) implementationUseCase.AuthUseCase {
// 	return *implementationUseCase.NewAuthUseCase(log)
// 	// Не нужно передавать пустую DTO, use case должен работать с интерфейсами
// }

/* Инициализация контроллера */

func main() {

	// Инициализация логгера
	log := logger.InitLogger()

	// Инициализация use case
	authUseCase := implementationUseCase.NewAuthUseCase(*log)

	// инициализация контроллера
	contro := controllers.NewController(*authUseCase, *log)

	// Инициализация router
	routs := routing.SetupRouter(contro)

	// инициализурем сервер
	server := &http.Server{
		Addr:    ":8080",
		Handler: routs,
	}

	log.PtintInfo(context.Background(), "Стартуем сервер на порте :8080")
	if err := server.ListenAndServe(); err != nil {
		log.PrintError(context.TODO(), "Сервер не стартонул", err)
	}

	// // маршрут с middleware
	// server.engin.GET("/secure", AuthMidlleware(), gin.WrapH(http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Secure Context"))
	// }))))

	// // обычный маршрут

	// server.RegisterPublicRoute(http.MethodGet, "/app", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Public context"))
	// })

	// if err := server.Start(":8080"); err != nil {
	// 	panic(err)
	// }

}
