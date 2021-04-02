package forms

import (
	"time"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Building struct {
	ID uint `gorm:"primaryKey" param:"ID" query:"ID" json:"ID"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string `param:"Name" query:"Name" json:"Name"`
	Description string `param:"Description" query:"Description" json:"Description"`
	OpeningDay string `param:"OpeningDay" query:"OpeningDay" json:"OpeningDay"`
	OpeningHour string `param:"OpeningHour" query:"OpeningHour" json:"OpeningHour"`
	Address string `param:"Address" query:"Address" json:"Address"`
	Phone string `param:"Phone" query:"Phone" json:"Phone"`
	Transport datatypes.JSON `param:"Transport" query:"Transport" json:"Transport"`
	Coordinates datatypes.JSON `param:"Coordinates" query:"Coordinates" json:"Coordinates"`
	Image string `param:"Image" query:"Image" json:"Image"`
	Note string `param:"Note" query:"Note" json:"Note"`
	Equipment string `param:"Equipment" query:"Equipment" json:"Equipment"`
	Rooms datatypes.JSON `param:"Rooms" query:"Rooms" json:"Rooms"`
}