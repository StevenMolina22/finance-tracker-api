package controller

import "github.com/gofiber/fiber/v2"

func (ctlr *Controller) GetRoot(c *fiber.Ctx) error {
	return c.SendString("Server active")
}
