package models

import (
	"echo-demo/db"
	dto "echo-demo/DTOs"
)

func CreateUser(user dto.User) (dto.User, error) {
	err := db.DB.Create(&user).Error
	return user, err
}

func UpdateUser(user dto.User) (dto.User, error) {
	err := db.DB.Save(&user).Error
	return user, err
}

func GetUserById(id uint) (user dto.User, err error) {
	err = db.DB.Table("users").First(&user, id).Error
	return user, err
}

func GetUserByUsername(username string) (user dto.User, err error) {
	err = db.DB.Table("users").Where("username = ?", username).First(&user).Error
	return user, err
}
