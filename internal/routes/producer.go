package routes

import (
	"book-explorer-es/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupProducerRoutes(route fiber.Router) {
	// Initialize controller
	producerController := controllers.NewProducerController()

	// Producer routes
	route.Post("/producers", producerController.CreateProducer)
	route.Get("/producers/:id", producerController.GetProducer)
	route.Get("/producers", producerController.GetAllProducers)
	route.Put("/producers/:id", producerController.UpdateProducer)
	route.Delete("/producers/:id", producerController.DeleteProducer)
}
