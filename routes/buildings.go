package routes

import (
	"github.com/labstack/echo/v4"
	"echo-demo/controllers"
)

func BuildingsRoute(e *echo.Group) {
	buildings := e.Group("/buildings")
	buildings.GET("", controllers.GetBuildings)
	buildings.GET("/:id/rooms", controllers.GetRoomsByBuildingId)
}