package routes

import (
	"book-explorer-es/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCarRoutes(route fiber.Router) {
	// Initialize controller
	carController := controllers.NewCarController()

	// Car routes
	route.Post("/cars", carController.CreateCar)
	route.Get("/cars/:id", carController.GetCar)
	route.Get("/cars", carController.GetAllCars)
	route.Put("/cars/:id", carController.UpdateCar)
	route.Delete("/cars/:id", carController.DeleteCar)
}
