package models

import (
	"gorm.io/gorm"

	database "echo-demo/db"
	form "echo-demo/forms"
)

func GetBuildings() (buildings []form.Building, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		err = db.Model(&buildings).Find(&buildings).Error
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
									array_to_json(array(select json_build_object('ID', id, 'CreatedAt',created_at,
																							'UpdatedAt', updated_at, 'DeletedAt', deleted_at,
																							'Name', name, 'Description', description,
																							'NumberOfPeople', number_of_people,
																							'Area', area, 'BuildingID', building_id)
																			from rooms where building_id = 1))
										as rooms
									from buildings`).Scan(&buildings).Error
	}
	return buildings, err
}