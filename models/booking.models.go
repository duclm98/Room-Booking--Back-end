package models

import (
	"echo-demo/db"
	dto "echo-demo/DTOs"
)

func GetBookingByBuildingIdDateTime(buildingID uint, date string, time string) (bookingArr []dto.Booking, err error) {
	err = db.DB.Table("booking").Where("building_id = ? AND date = ? AND time = ?", buildingID, date, time).Find(&bookingArr).Error
	return bookingArr, err
}
