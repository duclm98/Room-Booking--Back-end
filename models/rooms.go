package models

import (

	"gorm.io/gorm"

	database "echo-demo/db"
	form "echo-demo/forms"
)

func GetRoomsByBuildingId(id uint64) (rooms []form.Room, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		db.Table("rooms").Where("building_id = ?", id).Find(&rooms)
	}
	return rooms, err
}