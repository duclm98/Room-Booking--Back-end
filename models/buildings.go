package models

import (
	"echo-demo/db"
)

func GetBuildings() (buildings []Building, err error) {
	db, err := db.Connect()
	if err == nil {
		db.Table("buildings").Find(&buildings)
	}
	return buildings, err
}

func GetBuildings2() (buildings []Building, err error) {
	db, err := db.Connect()
	if err == nil {
		err = db.Table("buildings").Find(&buildings).Error
		for i := range buildings {
			db.Table("rooms").Where("buildingid = ?", buildings[i].ID).Find(&buildings[i].Rooms)
		}
	}
	return buildings, err
}