package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	form "echo-demo/forms"
	"echo-demo/models"
)

func GetBuildings(c echo.Context) error {
	building, err := models.GetBuildings()
	// building, err := models.GetBuildings2()
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, building)
}

func GetRoomsByBuildingId(c echo.Context) error {
	r := new(form.Room)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	rooms, err := models.GetRoomsByBuildingId(r.BuildingID)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, rooms)
}

func GetAvailableRooms(c echo.Context) error {
	b := new(form.Booking)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingArr, err := models.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err)
	}

	roomsArr, err := models.GetRoomsByBuildingId(b.BuildingID)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err)
	}

	var result []form.Room

	for i := range roomsArr {
		if a := contains(bookingArr, roomsArr[i].ID); a == false {
			result = append(result, roomsArr[i])
		}
	}

	return c.JSON(http.StatusOK, result)
}

func contains(arr []form.Booking, id uint) bool {
	for i := range arr {
		if arr[i].RoomID == id {
			return true
		}
	}
	return false
}