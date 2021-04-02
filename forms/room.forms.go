package forms

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name           string
	Description    string
	NumberOfPeople uint
	Area           uint
	BuildingID 		 uint
}