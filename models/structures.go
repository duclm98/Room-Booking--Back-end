package models

import (
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type Building struct {
	gorm.Model
	Name string
	Description string
	OpeningDay string
	OpeningHour string
	Address string
	Phone string
	Transport datatypes.JSON
	Coordinates datatypes.JSON
	Image string
	Note string
	Equipment string
	Rooms []Room
}

type Room struct {
	gorm.Model
	Name           string
	Description    string
	NumberOfPeople uint
	Area           uint
	BuildingID 		 uint
}