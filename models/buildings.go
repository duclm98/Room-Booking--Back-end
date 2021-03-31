package models

import (
	"fmt"

	"gorm.io/gorm"
	
	"echo-demo/db"
)

type Transport struct {
	Bus []string
	Oto []string
}

type Coordinates struct {
	Longitude string
	Latitude string
}

type Building struct {
	gorm.Model
	Name string
	Description string
	OpeningDay string
	OpeningHour string
	Address string
	Phone string
	Transport Transport
	Coordinates Coordinates
	Image string
	Note string
	Equipment string
}

func GetBuildings() {
	db, err := db.Connect()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var building []Building
	db.Table("buildings").Find(&building)
	fmt.Println(building)
}