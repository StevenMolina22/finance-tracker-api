package authors

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type AuthorHandler struct {
	queries *Queries
}

func NewHandler(db *sql.DB) *AuthorHandler {
	return &AuthorHandler{
		queries: New(db),
	}
}

func (hdlr *AuthorHandler) GetAuthors(c *fiber.Ctx) error {
	ctx := c.Context()

	authors, err := hdlr.queries.ListAuthors(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(authors)
}

func (hdlr *AuthorHandler) CreateAuthor(c *fiber.Ctx) error {
	ctx := c.Context()

	// Parse request body
	var rawAuthor CreateAuthorParams
	if err := c.BodyParser(&rawAuthor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create author
	author, err := hdlr.queries.CreateAuthor(ctx, rawAuthor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}
