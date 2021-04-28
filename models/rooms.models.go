package models

import (
	"echo-demo/db"
	dto "echo-demo/DTOs"
)

func GetRoomsByBuildingId(BuildingID uint) (rooms []dto.Room, err error) {
	db.DB.Table("rooms").Where("building_id = ?", BuildingID).Find(&rooms)
	return rooms, err
}
