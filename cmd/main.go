package main

import (
	"book-explorer-es/internal/database"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// listen to http request
	port := viper.Get("PORT")
	if port == nil {
		port = "8080"
	}
	app.Listen(fmt.Sprintf(":%s", port))
}
