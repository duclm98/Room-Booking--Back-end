package forms

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Date string
	Time string
	BuildingID uint
	RoomID uint
}