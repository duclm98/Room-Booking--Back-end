package forms

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID uint `gorm:"primaryKey" param:"ID" query:"ID" json:"ID"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Date string `param:"Date" query:"Date" json:"Date"`
	Time string `param:"Time" query:"Time" json:"Time"`
	BuildingID uint `param:"BuildingID" query:"BuildingID" json:"BuildingID"`
	RoomID uint `param:"RoomID" query:"RoomID" json:"RoomID"`
}