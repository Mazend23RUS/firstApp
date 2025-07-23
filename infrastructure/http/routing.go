package server

import (
	"github.com/alexey/firstApp/adapters/controllers"
	serverinterface "github.com/alexey/firstApp/infrastructure/restServer"
)

type Routing interface {
	RoutingServer()
}

func (rout *Rout) RoutingServer() {
	rout.RegisterPublicRoute()
}

type Rout struct {
	server     serverinterface.Server
	controller controllers.UserController
}

func SetupRouter(server serverinterface.Server, contr controllers.UserController) *Rout {
	router := &Rout{
		controller: contr,
		server:     server,
	}
	router.RoutingServer() /* Вызываем регистрацию маршрутов */
	return router

}

func (rout *Rout) RegisterPublicRoute() {

	rout.server.RegisterPublicRoute("POST", "/auth/user", rout.controller.LoginHandler)
	rout.server.RegisterPublicRoute("POST", "/auth/botton", rout.controller.ButtonHandler)

}
