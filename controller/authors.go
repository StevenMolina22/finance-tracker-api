package controller

import (
	"database/sql"

	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

// Implemented the controller so that the funcs use the same db queries instance
type Controller struct {
	queries *sqlc.Queries
}

// Constructor like func
func NewController(db *sql.DB) *Controller {
	return &Controller{
		queries: sqlc.New(db),
	}
}

func (ctlr *Controller) GetAuthors(c *fiber.Ctx) error {
	ctx := c.Context()

	authors, err := ctlr.queries.ListAuthors(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(authors)
}

func (ctlr *Controller) CreateAuthor(c *fiber.Ctx) error {
	ctx := c.Context()

	// Parse request body
	var rawAuthor sqlc.CreateAuthorParams
	if err := c.BodyParser(&rawAuthor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create author
	author, err := ctlr.queries.CreateAuthor(ctx, rawAuthor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}
