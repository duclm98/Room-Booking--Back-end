package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	dto "echo-demo/DTOs"
	method "echo-demo/methods"
	"echo-demo/models"
)

func GetBuildingsList(c echo.Context) error {
	// a := c.Get("user").(form.User) // Ép kiểu từ interface{} sang 1 interface khác
	// fmt.Println(a.Username)

	// building, err := models.GetBuildingsList()
	building, err := models.GetBuildingsList2()
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, building)
}

func GetBuilding(c echo.Context) error {
	b := new(dto.Building)
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
	r := new(dto.Room)
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
	b := new(dto.Booking)
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

	var result []dto.Room

	for i := range roomsArr {
		if a := method.Contains(bookingArr, roomsArr[i].ID); a == false {
			result = append(result, roomsArr[i])
		}
	}

	return c.JSON(http.StatusOK, result)
}