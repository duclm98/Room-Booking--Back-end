package forms

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID uint `gorm:"primaryKey" param:"id" query:"id" json:"id"`
  CreatedAt time.Time `param:"createdAt" query:"createdAt" json:"createdAt"`
  UpdatedAt time.Time	`param:"updatedAt" query:"updatedAt" json:"updatedAt"`
  DeletedAt gorm.DeletedAt `gorm:"index" param:"deletedAt" query:"deletedAt" json:"deletedAt"`
	Date string `param:"date" query:"date" json:"date"`
	Time string `param:"time" query:"time" json:"time"`
	BuildingID uint `param:"buildingId" query:"buildingId" json:"buildingId"`
	RoomID uint `param:"roomId" query:"roomId" json:"roomId"`
}