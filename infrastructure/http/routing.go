package routing

import (
	"github.com/alexey/adapters/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.UserController) *gin.Engine {

	rout := gin.New()

	// Глобальные middleware
	rout.Use(gin.Logger())
	rout.Use(gin.Recovery())

	// Public routes
	public := rout.Group("/api/ver")
	{
		public.POST("/login", gin.WrapF(authController.Logger))
		public.GET("/heep", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "ok"})
		})

	}

	// Protected routes (with auth middleware)
	// protected := rout.Group("/api/v1")
	// protected.Use()
	// {
	// 	protected.GET("/profile", gin.WrapF(authController.ProfileHandler))
	// 	protected.POST("/logout", gin.WrapF(authController.LogoutHandler))
	// }

	return rout

}
