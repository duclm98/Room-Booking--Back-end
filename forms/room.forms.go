package forms

import (
	"time"
	"gorm.io/gorm"
)

type Room struct {
	ID 						 uint 					`gorm:"primaryKey" param:"ID" query:"ID" json:"ID"`
  CreatedAt 	   time.Time
  UpdatedAt 	   time.Time
  DeletedAt 		 gorm.DeletedAt `gorm:"index"`
	Name           string 				`param:"Name" query:"Name" json:"Name"`
	Description    string  				`param:"Description" query:"Description" json:"Description"`
	NumberOfPeople uint 					`param:"NumberOfPeople" query:"NumberOfPeople" json:"NumberOfPeople"`
	Area           uint 					`param:"Area" query:"Area" json:"Area"`
	BuildingID 		 uint 					`param:"BuildingID" query:"BuildingID" json:"BuildingID"`
}