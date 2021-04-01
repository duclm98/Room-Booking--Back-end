package routes

import (
	"github.com/labstack/echo/v4"
	"echo-demo/controllers"
)

func BuildingsRoute(e *echo.Group) {
	router := e.Group("/buildings")
	router.GET("", controllers.GetBuildings)
	router.GET("/:id/rooms", controllers.GetRoomsByBuildingId)
}