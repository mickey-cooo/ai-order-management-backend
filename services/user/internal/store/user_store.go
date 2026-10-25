package store

import (
	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/schema"
	"gorm.io/gorm"
)

type UserStore struct{ db *gorm.DB }

func GetUserStore(db *gorm.DB) *UserStore { return &UserStore{db} }

func (s *UserStore) GetById(id string) (schema.User, error) {
	var user schema.User
	err := s.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (s *UserStore) GetByEmail(email string) (schema.User, error) {
	var user schema.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return user, err
}
