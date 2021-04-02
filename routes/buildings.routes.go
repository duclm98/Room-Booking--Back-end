package routes

import (
	"github.com/labstack/echo/v4"
	
	"echo-demo/controllers"
)

func BuildingsRoute(e *echo.Group) {
	router := e.Group("/buildings")
	router.GET("", controllers.GetBuildingsList)
	router.GET("/:ID", controllers.GetBuilding)
	router.GET("/:BuildingID/rooms", controllers.GetRoomsByBuildingId)
	router.GET("/:BuildingID/available-rooms-list", controllers.GetAvailableRoomsList)
	// router.GET("/:BuildingID/available-rooms-calendar", controllers.GetAvailableRoomsCalendar)
}