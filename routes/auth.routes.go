package routes

import (
	"github.com/labstack/echo/v4"
	
	controller "echo-demo/controllers"
)

func AuthRoute(e *echo.Group) {
	router := e.Group("/auth")
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}