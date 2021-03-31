package models

import (

	"gorm.io/gorm"

	database "echo-demo/db"
)

func GetRoomsByBuildingId(id uint64) (rooms []Room, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		db.Table("rooms").Where("buildingid = ?", id).Find(&rooms)
	}
	return rooms, err
}