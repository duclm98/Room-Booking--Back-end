package models

import (

	"gorm.io/gorm"

	database "echo-demo/db"
)

type Room struct {
	gorm.Model
	BuildingId 		 uint
	Name           string
	Description    string
	NumberOfPeople uint
	Area           uint
}

func GetRoomsByBuildingId(id uint64) (rooms []Room, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		db.Table("rooms").Find(&rooms)
	}
	return rooms, err
}