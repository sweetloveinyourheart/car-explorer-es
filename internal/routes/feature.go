package routes

import (
	"book-explorer-es/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupFeatureRoutes(route fiber.Router) {
	// Initialize controller
	featureController := controllers.NewFeatureController()

	// Feature routes
	route.Post("/features", featureController.CreateFeature)
	route.Get("/features/:id", featureController.GetFeature)
	route.Get("/features", featureController.GetAllFeatures)
	route.Put("/features/:id", featureController.UpdateFeature)
	route.Delete("/features/:id", featureController.DeleteFeature)
}
