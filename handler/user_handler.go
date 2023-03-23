package handler

import (
	"github.com/Wicho90/learn-fiber/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func User(c *fiber.Ctx) error {

	user := model.User{
		Name: "Albert",
		Age:  12,
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := model.NewUser()

	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.Id = uuid.NewString()

	return c.Status(fiber.StatusCreated).JSON(user)
}
