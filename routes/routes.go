package routes

import (
	"go-bored-api/controllers"
	"go-bored-api/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// public routes
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	config := middlewares.GetDefaultConfig()

	// protected routes
	admin := e.Group("/admin", echojwt.WithConfig(config.Init()))

	admin.GET("/users", controllers.GetAll)
	admin.DELETE("/users/:id", controllers.Delete)
}
