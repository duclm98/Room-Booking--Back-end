package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	dto "echo-demo/DTOs"
	"echo-demo/models"
)

func GetBookingByBuildingIdDateTime(c echo.Context) error {
	b := new(dto.Booking)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingArr, err := models.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookingArr)
}