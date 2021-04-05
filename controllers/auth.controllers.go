package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	form "echo-demo/forms"
	method "echo-demo/methods"
	model "echo-demo/models"
)

type Response struct {
	Message string `json:"message"`
}

func Register(c echo.Context) error {
	u := new(form.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := model.GetUserByUsername(u.Username)
	if err == nil {
		res := &Response{
			Message: "Tên đăng nhập đã tồn tại",
		}
		return c.JSON(http.StatusConflict, res)
	}

	hashPassword, err := method.HashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser := form.User {
		Username: u.Username,
		Password: hashPassword,
	}
	result, err := model.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func Login(c echo.Context) error {
	b := new(form.Booking)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingArr, err := model.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookingArr)
}