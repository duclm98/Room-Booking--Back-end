package models

import (
	"echo-demo/db"
	dto "echo-demo/DTOs"
)

func GetBuildingsList() (buildings []dto.Building, err error) {
	err = db.DB.Model(&buildings).Find(&buildings).Error
	return buildings, err
}

func GetBuildingsList2() (buildings []dto.Building, err error) {
	err = db.DB.Raw(`select *,
										array_to_json(array(select json_build_object('id', id, 'createdAt',created_at,
																								'updatedAt', updated_at, 'deletedAt', deleted_at,
																								'name', name, 'description', description,
																								'numberOfPeople', number_of_people,
																								'area', area, 'buildingId', building_id)
																				from rooms where building_id = buildings.id))
											as rooms
									from buildings`).Scan(&buildings).Error

	return buildings, err
}

func GetBuilding(id uint) (building dto.Building, err error) {
	err = db.DB.Model(&building).First(&building, id).Error
	return building, err
}
