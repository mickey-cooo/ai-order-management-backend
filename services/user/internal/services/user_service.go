package services

import (
	"context"

	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/dto"
	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/store"
)

type UserService struct {store *store.UserStore}


func NewUserService(s *store.UserStore) *UserService { return &UserService{s} }

func (s *UserService) GetUserById(ctx context.Context, id string) (dto.UserResponse, error) {
	u, err := s.store.GetById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.MapSchemaToResponse(u), nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	u, err := s.store.GetByEmail(email)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.MapSchemaToResponse(u), nil
}