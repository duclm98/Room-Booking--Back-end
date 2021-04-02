package models

import (

	"gorm.io/gorm"

	form "echo-demo/forms"
	database "echo-demo/db"
)

func GetBuildings() (buildings []form.Building, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		err = db.Table("buildings").Find(&buildings).Error
	}
	return buildings, err
}

func GetBuildings2() (buildings []form.Building, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		// err = db.Table("buildings").Find(&buildings).Error
		// if err == nil {
		// 	for i := range buildings {
		// 		db.Table("rooms").Where("buildingid = ?", buildings[i].ID).Find(&buildings[i].Rooms)
		// 	}
		// }
		err = db.Raw(`select *,
				array_to_json(array(select row_to_json(rooms) from rooms where building_id = buildings.id)) as rooms
				from buildings`).Scan(&buildings).Error
	}
	return buildings, err
}