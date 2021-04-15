package models

import (
	"gorm.io/gorm"

	database "echo-demo/db"
	form "echo-demo/forms"
)

func CreateUser(user form.User) (form.User, error) {
	var db *gorm.DB
	var err error
	db, err = database.Connect()
	if err == nil {
		err = db.Create(&user).Error
	}
	return user, err
}

func UpdateUser(user form.User) (form.User, error) {
	var db *gorm.DB
	var err error
	db, err = database.Connect()
	if err == nil {
		err = db.Save(&user).Error
	}
	return user, err
}

func GetUserById(id uint) (user form.User, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		err = db.Table("users").First(&user, id).Error
	}
	return user, err
}

func GetUserByUsername(username string) (user form.User, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		err = db.Table("users").Where("username = ?", username).First(&user).Error
	}
	return user, err
}