package models

import "gorm.io/gorm"

type OrganizationUser struct {
	gorm.Model
	UserID         uint `gorm:"primaryKey"`
	OrganizationID uint `gorm:"primaryKey"`
	IsAdmin        bool
}
