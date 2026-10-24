package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(50);uniqueIndex;not null"`
	Status      Status       `gorm:"type:varchar(20);default:'active';not null;check:chk_roles_status,status IN ('active','inactive','deleted','locked')"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
