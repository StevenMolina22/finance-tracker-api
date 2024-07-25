package router

import (
	"database/sql"

	"github.com/StevenMolina22/fiber-turso/controller"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(db *sql.DB, app *fiber.App) {
	ctlr := controller.NewController(db)

	app.Get("/", ctlr.GetRoot)

	// Authors
	authors := app.Group("/authors")
	authors.Get("/", ctlr.GetAuthors)
	authors.Post("/", ctlr.CreateAuthor)
}
