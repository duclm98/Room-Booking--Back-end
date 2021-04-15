package routes

import (

	"github.com/labstack/echo/v4"

	controller "echo-demo/controllers"
)

func BuildingsRoute(e *echo.Group) {
	router := e.Group("/buildings")
	router.GET("", controller.GetBuildingsList)
	router.GET("/:ID", controller.GetBuilding)
	router.GET("/:BuildingID/rooms", controller.GetRoomsByBuildingId)
	router.GET("/:BuildingID/available-rooms-list", controller.GetAvailableRoomsList)
	// router.GET("/:BuildingID/available-rooms-calendar", controller.GetAvailableRoomsCalendar)
}