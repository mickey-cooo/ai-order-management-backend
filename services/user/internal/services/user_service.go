package services

import (
	"context"

	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/dto"
	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/store"
)

type UserService struct{ store *store.UserStore }

func NewUserService(s *store.UserStore) *UserService { return &UserService{s} }

func (s *UserService) GetUserById(ctx context.Context, id string) (dto.UserResponse, error) {
	u, err := s.store.GetById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.MapSchemaToResponse(u), nil
}

func (s *UserService) GetUserListByIds(ctx context.Context, ids []string) ([]dto.UserResponse, error) {
	users, err := s.store.GetListByIds(ids)
	if err != nil {
		return nil, err
	}
	res := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, dto.MapSchemaToResponse(u))
	}
	return res, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	u, err := s.store.GetByEmail(email)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.MapSchemaToResponse(u), nil
}
