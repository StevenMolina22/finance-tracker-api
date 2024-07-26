package main

import (
	"log"
	"os"

	"github.com/StevenMolina22/fiber-turso/database"
	"github.com/StevenMolina22/fiber-turso/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	url := os.Getenv("DB_URL")

	// Open database
	db := database.Init(url)
	defer db.Close()

	app := fiber.New()
	router.PublicRoutes(db, app)
	log.Fatal(app.Listen(":3000"))
}
