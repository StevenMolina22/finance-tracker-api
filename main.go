package main

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/StevenMolina22/fiber-turso/authors"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//go:embed schemas.sql
var schemaSQL string

var db *sql.DB

func main() {
	var err error	
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := setupDatabase(); err != nil {
		log.Fatal(err)
	}

	app := setupFiber()
	log.Fatal(app.Listen(":4000"))
}

func setupFiber() *fiber.App {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Server active")
	})

	// Add more routes here
	app.Get("/authors", getAuthors)
	app.Post("/authors", createAuthor)

	return app
}

func setupDatabase() error {
	ctx := context.Background()

	// Create tables
	if _, err := db.ExecContext(ctx, schemaSQL); err != nil {
		return err
	}

	return nil
}

func getAuthors(c *fiber.Ctx) error {
	queries := authors.New(db)
	ctx := c.Context()

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(authors)
}

func createAuthor(c *fiber.Ctx) error {
	queries := authors.New(db)
	ctx := c.Context()

	// Parse request body
	var input authors.CreateAuthorParams
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create author
	author, err := queries.CreateAuthor(ctx, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}
