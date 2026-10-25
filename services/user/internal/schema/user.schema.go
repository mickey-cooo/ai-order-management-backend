package schema

type User struct {
	Base
	FirstName string  `gorm:"type:varchar(50)"`
	LastName  string  `gorm:"type:varchar(50)"`
	Email     string   `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  *string  `gorm:"type:varchar(255)"`
	Status    Status   `gorm:"type:varchar(20);default:'active';not null;check:chk_users_status,status IN ('active','inactive','deleted','locked')"`
	UserType  UserType `gorm:"type:varchar(20);default:'buyer';not null;check:chk_users_user_type,user_type IN ('buyer','supplier')"`
	Roles     []Role   `gorm:"many2many:user_roles"`
}
