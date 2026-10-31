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

func (s *UserStore) GetListByIds(ids []string) ([]schema.User, error) {
	var users []schema.User
	err := s.db.Where("id IN ?", ids).Find(&users).Error
	return users, err
}

func (s *UserStore) GetByEmail(email string) (schema.User, error) {
	var user schema.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return user, err
}
