package routes

import (
	"github.com/labstack/echo/v4"
	
	"echo-demo/controllers"
)

func AuthRoute(e *echo.Group) {
	router := e.Group("/auth")
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}