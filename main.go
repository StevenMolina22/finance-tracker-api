package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/StevenMolina22/fiber-turso/database"
	"github.com/StevenMolina22/fiber-turso/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Load env variables
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}
	url := os.Getenv("DB_URL")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Open database
	db := database.Init(url)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	app := fiber.New()

	router.PublicRoutes(db, app)
	log.Fatal(app.Listen(":" + port))
}
