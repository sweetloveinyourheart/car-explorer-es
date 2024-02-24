package controllers

import (
	"book-explorer-es/internal/models"
	"book-explorer-es/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CarController struct {
	svc services.ICarService
}

func NewCarController() *CarController {
	return &CarController{
		svc: services.NewCarService(),
	}
}

// CreateCar handles the creation of a new car
func (cmc *CarController) CreateCar(c *fiber.Ctx) error {
	car := new(models.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	if err := cmc.svc.CreateCar(car); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to create car",
		})
	}

	return c.JSON(car)
}

// GetCar retrieves a car by ID
func (pc *CarController) GetCar(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	car, err := pc.svc.GetCar(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Car not found",
		})
	}

	return c.JSON(car)
}

// GetAllCars retrieves all cars
func (pc *CarController) GetAllCars(c *fiber.Ctx) error {
	cars, err := pc.svc.GetAllCars()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to retrieve cars",
		})
	}

	return c.JSON(cars)
}

// UpdateCar updates a car
func (pc *CarController) UpdateCar(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	car := new(models.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	car.ID = uint(id)
	if err := pc.svc.UpdateCar(car); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to update car",
		})
	}

	return c.JSON(car)
}

// DeleteCar deletes a car by ID
func (pc *CarController) DeleteCar(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	if err := pc.svc.DeleteCar(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to delete car",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
