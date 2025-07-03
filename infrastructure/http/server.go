package server

import (
	"context"
	"fmt"
	"net/http"

	loggerinterface "github.com/alexey/pkg/logger/interface"
	"github.com/gin-gonic/gin"
)

/* GinServer - структура, которая инкапсулирует сервер Gin. Содержит поле engine - экземпляр Gin-движка. */
type GinServer struct {
	engin  *gin.Engine /* gin сслка на пакет import, вызываем у пакета структуру Engine с кучей различных методов */
	logger loggerinterface.Logger
}

/* NewGinServer() - конструктор, который:
Создаёт новый экземпляр Gin с помощью gin.New()
(в отличие от gin.Default(),
который автоматически добавляет Logger и Recovery middleware) */

func NewGinServer(log loggerinterface.Logger) *GinServer {
	server := &GinServer{
		engin:  gin.New(),
		logger: log,
	}

	/* Добавляем глобальные middleware */
	/* Явно добавляет два стандартных middleware: */
	server.engin.Use(gin.Logger())   /* Сервер логирует все действия */
	server.engin.Use(gin.Recovery()) /* Предотвращает падение сервера, отправляет ошибку 500 */

	return server
}

/*
	RegisterPublicRoute() - регистрирует обработчик для конкретного HTTP-метода и пути:

method - HTTP-метод (GET, POST и т.д.)
path - URL-путь
handler - стандартный обработчик http.HandlerFunc
gin.WrapH() адаптирует стандартный http.HandlerFunc к формату Gin
*/
func (g *GinServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	g.engin.Handle(method, path, gin.WrapF(handler))
}

/* Start() - запускает сервер на указанном адресе (например, ":8080") */
func (g *GinServer) Start(address string) error {
	logger := g.logger
	logger.PrintInfo(context.Background(), "Сервер стартанул на порту "+address)
	err := g.engin.Run(address)
	if err != nil {
		logger.PrintInfo(context.Background(), err.Error())
		return fmt.Errorf("Сервер не стартанул на порту "+address, err)
	}
	return nil
}

/* AuthMiddleware() - фабрика middleware, возвращающая функцию-обработчик типа gin.HandlerFunc */

func AuthMidlleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		if g.GetHeader("Authorization") == "" { /* Проверяется наличие заголовка Authorization */
			g.AbortWithStatus(http.StatusUnauthorized) /* Если заголовка нет - возвращается статус 401 (Unauthorized) */
			return                                     /* и обработка прерывается (AbortWithStatus) */

		}
		g.Next() /* Если заголовок есть - вызывается c.Next() для передачи управления следующему обработчику в цепочке */
	}
	/* Middleware в Gin - это функции, которые получают контекст запроса и могут:
	   Модифицировать запрос
	   Проверять условия
	   Прерывать выполнение
	   Передавать управление следующему обработчику */
}
