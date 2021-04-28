package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	config "echo-demo/config"
	controller "echo-demo/controllers"
	method "echo-demo/methods"
	model "echo-demo/models"
)

var env config.Config

func init() {
	var err error
	env, err = config.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load auth config:", err)
		return
	}
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := method.ExtractToken(c.Request())
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Không tìm thấy access token!",
			})
		}

		id, err := method.VerifyToken(token, env.AuthAccessTokenSecret)
		if err != nil || id < 1 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Access token không hợp lệ hoặc đã hết hạn!",
			})
		}

		user, err := model.GetUserById(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "User không tồn tại!",
			})
		}

		c.Set("user", user)

		return next(c)
	}
}

func Router(e *echo.Group) {
	authRouter := e.Group("/auth")
	authRouter.POST("/register", controller.Register)
	authRouter.POST("/login", controller.Login)
	authRouter.POST("/refresh-token", controller.Refresh)

	// buildingRouter := e.Group("/buildings", Authentication)
	buildingRouter := e.Group("/buildings")
	buildingRouter.GET("", controller.GetBuildingsList)
	buildingRouter.GET("/:ID", controller.GetBuilding)
	buildingRouter.GET("/:BuildingID/rooms", controller.GetRoomsByBuildingId)
	buildingRouter.GET("/:BuildingID/available-rooms-list", controller.GetAvailableRoomsList)

	bookingRouter := e.Group("/booking")
	bookingRouter.GET("", controller.GetBookingByBuildingIdDateTime)
}
