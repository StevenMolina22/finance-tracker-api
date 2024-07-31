package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctlr *Controller) ListBudgets(c *fiber.Ctx) error {
	budgets, err := ctlr.queries.ListBudgets(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(budgets)
}

func (ctlr *Controller) CreateBudget(c *fiber.Ctx) error {
	var rawBudget sqlc.CreateBudgetParams
	if err := c.BodyParser(&rawBudget); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	budget, err := ctlr.queries.CreateBudget(c.Context(), rawBudget)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(budget)
}

func (ctlr *Controller) GetBudget(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	budget, err := ctlr.queries.GetBudget(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(budget)
}

func (ctlr *Controller) UpdateBudget(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var rawBudget sqlc.UpdateBudgetParams
	if err := c.BodyParser(&rawBudget); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	rawBudget.ID = int64(id)

	budget, err := ctlr.queries.UpdateBudget(c.Context(), rawBudget)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(budget)
}

func (ctlr *Controller) DeleteBudget(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	budget, err := ctlr.queries.DeleteBudget(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(budget)
}
