package schema

import "time"

// Replaces gorm.Model: no DeletedAt (soft delete uses Status), gen_random_uuid needs PG 13+.
type Base struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	CreatedBy *string `gorm:"type:uuid"`
	UpdatedAt time.Time
	UpdatedBy *string `gorm:"type:uuid"`
}
