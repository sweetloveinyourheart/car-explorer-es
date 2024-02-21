package main

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	// read configs
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// init database connection
	database.InitPostgreSQLConnection()
	database.Migrate()

	// create fiber instances
	app := fiber.New()

	v1Router := app.Group("/api/v1")
	routes.SetupFeatureRoutes(v1Router)

	// listen to http request
	port := viper.Get("PORT")
	if port == nil {
		port = "8080"
	}
	app.Listen(fmt.Sprintf(":%s", port))
}
