package main

import (
	"github.com/Wicho90/learn-fiber/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())

	app.Get("/", handler.HelloWorld)

	userGroup := app.Group("/users")
	userGroup.Get("", handler.User)
	userGroup.Post("", handler.CreateUser)

	app.Listen(":8080")
}
