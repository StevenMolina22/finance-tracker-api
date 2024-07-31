package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctlr *Controller) GetTransactions(c *fiber.Ctx) error {
	ctx := c.Context()

	transactions, err := ctlr.queries.ListTransactions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(transactions)
}

func (ctlr *Controller) CreateTransaction(c *fiber.Ctx) error {
	ctx := c.Context()

	// Parse request body
	var rawTransaction sqlc.CreateTransactionParams
	if err := c.BodyParser(&rawTransaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create transaction
	transaction, err := ctlr.queries.CreateTransaction(ctx, rawTransaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(transaction)
}

func (ctlr *Controller) GetTransaction(c *fiber.Ctx) error {
	// get id param
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	transaction, err := ctlr.queries.GetTransaction(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(transaction)
}

func (ctlr *Controller) UpdateTransaction(c *fiber.Ctx) error {
	ctx := c.Context()

	// get id param
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Parse request body
	var rawTransaction sqlc.UpdateTransactionParams
	if err := c.BodyParser(&rawTransaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	rawTransaction.ID = int64(id)

	// Update transaction
	transaction, err := ctlr.queries.UpdateTransaction(ctx, rawTransaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(transaction)
}

func (ctlr *Controller) DeleteTransaction(c *fiber.Ctx) error {
	// get id param
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	transaction, err := ctlr.queries.DeleteTransaction(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(transaction)
}
