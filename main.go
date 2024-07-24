package main

import (
	"database/sql"
	_ "embed"
	"log"
	"os"

	"github.com/StevenMolina22/fiber-turso/authors"
	"github.com/StevenMolina22/fiber-turso/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	app := fiber.New()

	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dbUrl := os.Getenv("DB_URL")

	// Open database
	db := database.InitDB(dbUrl)
	defer db.Close()

	setupFiber(db, app)
	log.Fatal(app.Listen(":3000"))
}

func setupFiber(db *sql.DB, app *fiber.App) {

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Server active")
	})

	authorHdlr := authors.NewHandler(db)

	app.Get("/authors", authorHdlr.GetAuthors)
	app.Post("/authors", authorHdlr.CreateAuthor)
}
