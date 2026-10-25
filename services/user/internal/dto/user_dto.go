package dto

import (
	"time"

	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/schema"
)

type UserRequestById struct {
	Id string `json:"id" validate:"required"`
}

type UserResponse struct {
	Id        string      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapSchemaToResponse(user schema.User) UserResponse {
	return UserResponse{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}