package controllers

import (
	"book-explorer-es/internal/models"
	"book-explorer-es/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CarModelController struct {
	svc services.ICarModelService
}

func NewCarModelController() *CarModelController {
	return &CarModelController{
		svc: services.NewCarModelService(),
	}
}

// CreateCarModel handles the creation of a new car model
func (cmc *CarModelController) CreateCarModel(c *fiber.Ctx) error {
	carModel := new(models.CarModel)
	if err := c.BodyParser(carModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	if err := cmc.svc.CreateCarModel(carModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to create carModel",
		})
	}

	return c.JSON(carModel)
}

// GetCarModel retrieves a carModel by ID
func (pc *CarModelController) GetCarModel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	carModel, err := pc.svc.GetCarModel(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "CarModel not found",
		})
	}

	return c.JSON(carModel)
}

// GetAllCarModels retrieves all carModels
func (pc *CarModelController) GetAllCarModels(c *fiber.Ctx) error {
	carModels, err := pc.svc.GetAllCarModels()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to retrieve carModels",
		})
	}

	return c.JSON(carModels)
}

// UpdateCarModel updates a carModel
func (pc *CarModelController) UpdateCarModel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	carModel := new(models.CarModel)
	if err := c.BodyParser(carModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	carModel.ID = uint(id)
	if err := pc.svc.UpdateCarModel(carModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to update carModel",
		})
	}

	return c.JSON(carModel)
}

// DeleteCarModel deletes a carModel by ID
func (pc *CarModelController) DeleteCarModel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	if err := pc.svc.DeleteCarModel(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to delete carModel",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
