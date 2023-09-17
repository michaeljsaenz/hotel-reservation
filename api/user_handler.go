package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaeljsaenz/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "mike",
		LastName:  "boom",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("mike")
}
