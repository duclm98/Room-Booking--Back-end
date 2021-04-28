package DTOs

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID             uint           `gorm:"primaryKey" param:"id" query:"id" json:"id"`
	CreatedAt      time.Time      `param:"createdAt" query:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time      `param:"updatedAt" query:"updatedAt" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" param:"deletedAt" query:"deletedAt" json:"deletedAt"`
	Name           string         `param:"name" query:"name" json:"name"`
	Description    string         `param:"description" query:"description" json:"description"`
	NumberOfPeople uint           `param:"numberOfPeople" query:"numberOfPeople" json:"numberOfPeople"`
	Area           uint           `param:"area" query:"area" json:"area"`
	BuildingID     uint           `param:"buildingId" query:"buildingId" json:"buildingId"`
}
