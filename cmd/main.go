package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	// read configs
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// create fiber instances
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// listen to http request
	port := viper.Get("PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
