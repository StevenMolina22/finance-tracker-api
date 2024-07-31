package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctlr *Controller) ListUsers(c *fiber.Ctx) error {
	ctx := c.Context()

	users, err := ctlr.queries.ListUsers(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error fetching users: ": err.Error(),
		})
	}

	return c.JSON(users)
}

func (ctlr *Controller) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()

	// Parse request body
	var rawUser sqlc.CreateUserParams
	if err := c.BodyParser(&rawUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error parsing id: ": err.Error(),
		})
	}

	// Create user
	user, err := ctlr.queries.CreateUser(ctx, rawUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (ctlr *Controller) GetUser(c *fiber.Ctx) error {
	// get id param
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := ctlr.queries.GetUser(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error fetching user: ": err.Error(),
		})
	}
	return c.JSON(user)
}

func (cltr *Controller) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Incorrect id format: ": err.Error(),
		})
	}

	var rawUser sqlc.UpdateUserParams
	if err := c.BodyParser(&rawUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error updating user: ": err.Error(),
		})
	}
	rawUser.ID = int64(id) // set id by params and not by body

	user, err := cltr.queries.UpdateUser(c.Context(), rawUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error updating user: ": err.Error(),
		})
	}

	return c.JSON(user)
}

func (ctlr *Controller) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error parsing id: ": err.Error(),
		})
	}

	user, err := ctlr.queries.DeleteUser(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error deleting user: ": err.Error(),
		})
	}

	return c.JSON(user)
}
