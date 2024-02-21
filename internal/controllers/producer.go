package controllers

import (
	"book-explorer-es/internal/models"
	"book-explorer-es/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProducerController struct {
	svc services.IProducerService
}

func NewProducerController() *ProducerController {
	return &ProducerController{
		svc: services.NewProducerService(),
	}
}

// CreateProducer handles the creation of a new producer
func (pc *ProducerController) CreateProducer(c *fiber.Ctx) error {
	producer := new(models.Producer)
	if err := c.BodyParser(producer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	if err := pc.svc.CreateProducer(producer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to create producer",
		})
	}

	return c.JSON(producer)
}

// GetProducer retrieves a producer by ID
func (pc *ProducerController) GetProducer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	producer, err := pc.svc.GetProducer(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Producer not found",
		})
	}

	return c.JSON(producer)
}

// GetAllProducers retrieves all producers
func (pc *ProducerController) GetAllProducers(c *fiber.Ctx) error {
	producers, err := pc.svc.GetAllProducers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to retrieve producers",
		})
	}

	return c.JSON(producers)
}

// UpdateProducer updates a producer
func (pc *ProducerController) UpdateProducer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	producer := new(models.Producer)
	if err := c.BodyParser(producer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	producer.ID = uint(id)
	if err := pc.svc.UpdateProducer(producer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to update producer",
		})
	}

	return c.JSON(producer)
}

// DeleteProducer deletes a producer by ID
func (pc *ProducerController) DeleteProducer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	if err := pc.svc.DeleteProducer(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to delete producer",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
