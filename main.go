package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"os"

	"github.com/StevenMolina22/fiber-turso/authors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type AuthorHandler struct {
	queries *authors.Queries
}

func NewAuthorHandler(db *sql.DB) *AuthorHandler {
	return &AuthorHandler{
		queries: authors.New(db),
	}
}

//go:embed schemas.sql
var schemaSQL string

func main() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dbUrl := os.Getenv("DB_URL")

	// Open database
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Create tables
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, schemaSQL); err != nil {
		log.Fatal("Error creating tables: ", err)
	}

	app := setupFiber(db)
	log.Fatal(app.Listen(":3000"))
}

func setupFiber(db *sql.DB) *fiber.App {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Server active")
	})

	authorHdlr := NewAuthorHandler(db)

	app.Get("/authors", authorHdlr.getAuthors)
	app.Post("/authors", authorHdlr.createAuthor)

	return app
}
