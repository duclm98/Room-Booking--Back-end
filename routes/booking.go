package routes

import (
	"github.com/labstack/echo/v4"
	"echo-demo/controllers"
)

func BookingRoute(e *echo.Group) {
	router := e.Group("/booking")
	router.GET("", controllers.GetBookingByBuildingIdDateTime)
}