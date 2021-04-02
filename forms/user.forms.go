package forms

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID 						 uint 					`gorm:"primaryKey" param:"ID" query:"ID" json:"ID"`
  CreatedAt 	   time.Time
  UpdatedAt 	   time.Time
  DeletedAt 		 gorm.DeletedAt `gorm:"index"`
	Username       string 				`param:"Username" query:"Username" json:"Username"`
	Password       string  				`param:"Password" query:"Password" json:"Password"`
}