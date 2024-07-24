package main

import (
	"github.com/StevenMolina22/fiber-turso/authors"
	"github.com/gofiber/fiber/v2"
)

func (hdlr *AuthorHandler) getAuthors(c *fiber.Ctx) error {
	ctx := c.Context()

	authors, err := hdlr.queries.ListAuthors(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(authors)
}

func (hdlr *AuthorHandler) createAuthor(c *fiber.Ctx) error {
	ctx := c.Context()

	// Parse request body
	var rawAuthor authors.CreateAuthorParams
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
