package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) sendJSON(c *fiber.Ctx, status int, errMsg string, data interface{}) error {
	response := fiber.Map{
		"error": errMsg != "",
		"msg":   errMsg,
		"data":  data,
	}
	return c.Status(status).JSON(response)
}

func (ctr *Controller) ListGoals(c *fiber.Ctx) error {
	goals, err := ctr.queries.ListGoals(c.Context())
	if err != nil {
		return ctr.sendJSON(c, 500, "Failed to list goals", nil)
	}
	return ctr.sendJSON(c, 200, "", goals)
}

func (ctr *Controller) CreateGoal(c *fiber.Ctx) error {
	var rawGoal sqlc.CreateGoalParams
	if err := c.BodyParser(&rawGoal); err != nil {
		return ctr.sendJSON(c, 400, "Failed to parse request body", nil)
	}

	goal, err := ctr.queries.CreateGoal(c.Context(), rawGoal)
	if err != nil {
		return ctr.sendJSON(c, 500, "Failed to create goal", nil)
	}
	return ctr.sendJSON(c, 200, "", goal)
}

func (ctr *Controller) GetGoal(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ctr.sendJSON(c, 400, "Failed to parse goal ID", nil)
	}

	goal, err := ctr.queries.GetGoal(c.Context(), int64(id))
	if err != nil {
		return ctr.sendJSON(c, 500, "Failed to get goal", nil)
	}
	return ctr.sendJSON(c, 200, "", goal)
}

func (ctr *Controller) UpdateGoal(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ctr.sendJSON(c, 400, "Failed to parse goal ID", nil)
	}

	var rawGoal sqlc.UpdateGoalParams
	if err := c.BodyParser(&rawGoal); err != nil {
		return ctr.sendJSON(c, 400, "Failed to parse request body", nil)
	}
	rawGoal.ID = int64(id)

	goal, err := ctr.queries.UpdateGoal(c.Context(), rawGoal)
	if err != nil {
		return ctr.sendJSON(c, 500, "Failed to update goal", nil)
	}
	return ctr.sendJSON(c, 200, "", goal)
}

func (ctr *Controller) DeleteGoal(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ctr.sendJSON(c, 400, "Failed to parse goal ID", nil)
	}

	goal, err := ctr.queries.DeleteGoal(c.Context(), int64(id))
	if err != nil {
		return ctr.sendJSON(c, 500, "Failed to delete goal", nil)
	}
	return ctr.sendJSON(c, 200, "", goal)
}

func (ctr *Controller) ListGoalsByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return ctr.sendJSON(c, 400, "Failed to parse user ID", nil)
	}

	goals, err := ctr.queries.GetGoalsByUserId(c.Context(), int64(userID))
	if err != nil {
		return ctr.sendJSON(c, 500, "Failed to get goals", nil)
	}
	return ctr.sendJSON(c, 200, "", goals)
}
