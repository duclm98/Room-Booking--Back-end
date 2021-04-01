package models

import (

	"gorm.io/gorm"

	database "echo-demo/db"
	form "echo-demo/forms"
)

func GetBookingByBuildingIdDateTime(buildingID uint, date string, time string) (bookingArr []form.Booking, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		err = db.Table("booking").Where("building_id = ? AND date = ? AND time = ?", buildingID, date, time).Find(&bookingArr).Error
	}
	return bookingArr, err
}