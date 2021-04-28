package DTOs

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Building struct {
	ID          uint           `gorm:"primaryKey" param:"id" query:"id" json:"id"`
	CreatedAt   time.Time      `param:"createdAt" query:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time      `param:"updatedAt" query:"updatedAt" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" param:"deletedAt" query:"deletedAt" json:"deletedAt"`
	Name        string         `param:"name" query:"name" json:"name"`
	Description string         `param:"description" query:"description" json:"description"`
	OpeningDay  string         `param:"openingDay" query:"openingDay" json:"openingDay"`
	OpeningHour string         `param:"openingHour" query:"openingHour" json:"openingHour"`
	Address     string         `param:"address" query:"address" json:"address"`
	Phone       string         `param:"phone" query:"phone" json:"phone"`
	Transport   datatypes.JSON `param:"transport" query:"transport" json:"transport"`
	Coordinates datatypes.JSON `param:"coordinates" query:"coordinates" json:"coordinates"`
	Image       string         `param:"image" query:"image" json:"image"`
	Note        string         `param:"note" query:"note" json:"note"`
	Equipment   string         `param:"equipment" query:"equipment" json:"equipment"`
	Rooms       datatypes.JSON `param:"rooms" query:"rooms" json:"rooms"`
}
