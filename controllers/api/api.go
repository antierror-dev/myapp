package api_controller

import (
"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
return c.SendString("TestApi")
}