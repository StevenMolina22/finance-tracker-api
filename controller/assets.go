package controller

import (
	"github.com/StevenMolina22/fiber-turso/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (ctlr *Controller) ListAssets(c *fiber.Ctx) error {
	assets, err := ctlr.queries.ListAssets(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(assets)
}

func (ctlr *Controller) CreateAsset(c *fiber.Ctx) error {
	var rawAsset sqlc.CreateAssetParams
	if err := c.BodyParser(&rawAsset); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	asset, err := ctlr.queries.CreateAsset(c.Context(), rawAsset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(asset)
}

func (ctlr *Controller) GetAsset(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	asset, err := ctlr.queries.GetAsset(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(asset)
}

func (ctlr *Controller) UpdateAsset(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var rawAsset sqlc.UpdateAssetParams
	if err := c.BodyParser(&rawAsset); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	rawAsset.ID = int64(id)

	asset, err := ctlr.queries.UpdateAsset(c.Context(), rawAsset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(asset)
}

func (ctlr *Controller) DeleteAsset(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	asset, err := ctlr.queries.DeleteAsset(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(asset)
}
