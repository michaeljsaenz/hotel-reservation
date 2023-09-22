package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/michaeljsaenz/hotel-reservation/db"
	"github.com/michaeljsaenz/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore // interface
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "mike",
		LastName:  "boom",
	}
	return c.JSON(u)
}
