package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	form "echo-demo/forms"
	"echo-demo/models"
)

func Register(c echo.Context) error {
	b := new(form.Booking)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingArr, err := models.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookingArr)
}

func Login(c echo.Context) error {
	b := new(form.Booking)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingArr, err := models.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookingArr)
}