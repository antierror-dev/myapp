package controller


import (
  "github.com/gofiber/fiber/v2"
)


func Home(c *fiber.Ctx) error {
  return c.JSON(fiber.Map{
    "Hello":"This is home page",
  })
}
//c.Render("home",fiber.Map{"name":"MmdAli",})