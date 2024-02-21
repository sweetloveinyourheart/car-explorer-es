package controllers

import (
	"book-explorer-es/internal/models"
	"book-explorer-es/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FeatureController struct {
	svc services.IFeatureService
}

func NewFeatureController() *FeatureController {
	return &FeatureController{
		svc: services.NewFeatureService(),
	}
}

// CreateFeature handles the creation of a new feature
func (pc *FeatureController) CreateFeature(c *fiber.Ctx) error {
	feature := new(models.Feature)
	if err := c.BodyParser(feature); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	if err := pc.svc.CreateFeature(feature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to create feature",
		})
	}

	return c.JSON(feature)
}

// GetFeature retrieves a feature by ID
func (pc *FeatureController) GetFeature(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	feature, err := pc.svc.GetFeature(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Feature not found",
		})
	}

	return c.JSON(feature)
}

// GetAllFeatures retrieves all features
func (pc *FeatureController) GetAllFeatures(c *fiber.Ctx) error {
	features, err := pc.svc.GetAllFeatures()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to retrieve features",
		})
	}

	return c.JSON(features)
}

// UpdateFeature updates a feature
func (pc *FeatureController) UpdateFeature(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	feature := new(models.Feature)
	if err := c.BodyParser(feature); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	feature.ID = uint(id)
	if err := pc.svc.UpdateFeature(feature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to update feature",
		})
	}

	return c.JSON(feature)
}

// DeleteFeature deletes a feature by ID
func (pc *FeatureController) DeleteFeature(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	if err := pc.svc.DeleteFeature(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to delete feature",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
