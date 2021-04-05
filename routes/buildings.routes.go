package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	controller "echo-demo/controllers"
	myMiddleware "echo-demo/middlewares"
)

func IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuth := myMiddleware.Verify(c.Request())
		if isAuth == false {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Chưa đăng nhập vào hệ thống!",
			})
		}
		return next(c)
	}
}


var isAuth = myMiddleware.IsAuth

func BuildingsRoute(e *echo.Group) {
	router := e.Group("/buildings")
	router.GET("", controller.GetBuildingsList, IsAuth)
	router.GET("/:ID", controller.GetBuilding)
	router.GET("/:BuildingID/rooms", controller.GetRoomsByBuildingId)
	router.GET("/:BuildingID/available-rooms-list", controller.GetAvailableRoomsList)
	// router.GET("/:BuildingID/available-rooms-calendar", controller.GetAvailableRoomsCalendar)
}