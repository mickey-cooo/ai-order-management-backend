package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Roles []Role `gorm:"many2many:role_permissions"`
}
