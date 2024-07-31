package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctlr *Controller) ListCategories(c *fiber.Ctx) error {
	categories, err := ctlr.queries.ListCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(categories)
}

func (ctlr *Controller) CreateCategory(c *fiber.Ctx) error {
	var rawCategory sqlc.CreateCategoryParams
	if err := c.BodyParser(&rawCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := ctlr.queries.CreateCategory(c.Context(), rawCategory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

func (ctlr *Controller) GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := ctlr.queries.GetCategory(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(category)
}

func (ctlr *Controller) UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var rawCategory sqlc.UpdateCategoryParams
	if err := c.BodyParser(&rawCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rawCategory.ID = int64(id)
	category, err := ctlr.queries.UpdateCategory(c.Context(), rawCategory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(category)
}

func (ctlr *Controller) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := ctlr.queries.DeleteCategory(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(category)
}