package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctlr *Controller) ListLiabilities(c *fiber.Ctx) error {
	liabilities, err := ctlr.queries.ListLiabilities(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(liabilities)
}

func (ctlr *Controller) CreateLiability(c *fiber.Ctx) error {
	var rawLiability sqlc.CreateLiabilityParams
	if err := c.BodyParser(&rawLiability); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	liability, err := ctlr.queries.CreateLiability(c.Context(), rawLiability)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(liability)
}

func (ctlr *Controller) GetLiability(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	liability, err := ctlr.queries.GetLiability(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(liability)
}

func (ctlr *Controller) UpdateLiability(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var rawLiability sqlc.UpdateLiabilityParams
	if err := c.BodyParser(&rawLiability); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rawLiability.ID = int64(id)
	liability, err := ctlr.queries.UpdateLiability(c.Context(), rawLiability)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(liability)
}

func (ctlr *Controller) DeleteLiability(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	liability, err := ctlr.queries.DeleteLiability(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(liability)
}