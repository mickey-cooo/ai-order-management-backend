package handlers

import (
	"strings"

	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/services"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct{ svc *services.UserService }

func NewUserHandler(svc *services.UserService) *UserHandler { return &UserHandler{svc} }

// GetById godoc
//
//	@Summary		Get a user
//	@Description	Returns the user with the given UUID as {"status":200,"data":UserResponse}; a failure returns {"status":400,"error":"..."} with HTTP 200.
//	@Tags			users
//	@Produce		json
//	@Param			id	path		string	true	"user UUID"	Format(uuid)
//	@Success		200	{object}	map[string]interface{}
//	@Router			/users/{id} [get]
func (h *UserHandler) GetById(c fiber.Ctx) error {
	u, err := h.svc.GetUserById(c.Context(), c.Params("id"))

	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   u,
	})
}

func (h *UserHandler) GetListByIds(c fiber.Ctx) error {
	// FieldsFunc drops empty parts so "?ids=" yields no ids, not [""]
	ids := strings.FieldsFunc(c.Query("ids"), func(r rune) bool { return r == ',' })
	users, err := h.svc.GetUserListByIds(c.Context(), ids)

	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   users,
	})
}
