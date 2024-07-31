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

func (ctlr *Controller) GetRoot(c *fiber.Ctx) error {
	return c.SendString("Server active")
}