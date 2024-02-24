package routes

import (
	"book-explorer-es/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCarModelRoutes(route fiber.Router) {
	// Initialize controller
	carModelController := controllers.NewCarModelController()

	// CarModel routes
	route.Post("/carModels", carModelController.CreateCarModel)
	route.Get("/carModels/:id", carModelController.GetCarModel)
	route.Get("/carModels", carModelController.GetAllCarModels)
	route.Put("/carModels/:id", carModelController.UpdateCarModel)
	route.Delete("/carModels/:id", carModelController.DeleteCarModel)
}
