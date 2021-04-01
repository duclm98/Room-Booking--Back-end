package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	
	"echo-demo/models"
)

type Booking struct {
	ID uint `param:"ID" query:"ID" json:"ID"`
	BuildingID uint `param:"BuildingID" query:"BuildingID" json:"BuildingID"`
	RoomID uint `param:"RoomID" query:"RoomID" json:"RoomID"`
	Date string `param:"Date" query:"Date" json:"Date"`
	Time string `param:"Time" query:"Time" json:"Time"`
}

func GetBookingByBuildingIdDateTime(c echo.Context) error {
	b := new(Booking)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingArr, err := models.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookingArr)
}