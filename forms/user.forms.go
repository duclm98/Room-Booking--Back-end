package forms

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" param:"id" query:"id" json:"id"`
	CreatedAt    time.Time      `param:"createdAt" query:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time      `param:"updatedAt" query:"updatedAt" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" param:"deletedAt" query:"deletedAt" json:"deletedAt"`
	Username     string         `param:"username" query:"username" json:"username"`
	Password     string         `param:"password" query:"password" json:"password"`
	RefreshToken string         `param:"refreshToken" query:"refreshToken" json:"refreshToken"`
}
