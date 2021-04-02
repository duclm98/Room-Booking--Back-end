package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"echo-demo/models"
	form "echo-demo/forms"
	method "echo-demo/methods"
)

func GetBuildingsList(c echo.Context) error {
	building, err := models.GetBuildingsList()
	// building, err := models.GetBuildingsList2()
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, building)
}

func GetBuilding(c echo.Context) error {
	b := new(form.Building)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	building, err := models.GetBuilding(b.ID)
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

func GetAvailableRoomsList(c echo.Context) error {
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
		if a := method.Contains(bookingArr, roomsArr[i].ID); a == false {
			result = append(result, roomsArr[i])
		}
	}

	return c.JSON(http.StatusOK, result)
}

// func GetAvailableRoomsCalendar(c echo.Context) error {
// 	b := new(form.Booking)
// 	if err := c.Bind(b); err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	bookingArr, err := models.GetBookingByBuildingIdDateTime(b.BuildingID, b.Date, b.Time)
// 	if(err != nil) {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	roomsArr, err := models.GetRoomsByBuildingId(b.BuildingID)
// 	if(err != nil) {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	timeArr := method.TimeArray()

// 	result := Result {
// 		TimeArray: timeArr
// 	}

// 	for i := range roomsArr {
// 		if a := method.Contains(bookingArr, roomsArr[i].ID); a == false {
			
// 		}
// 	}

// 	return c.JSON(http.StatusOK, result)
// }