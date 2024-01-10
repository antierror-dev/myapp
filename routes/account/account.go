package account_routes

import (
  "github.com/gofiber/fiber/v2"
  controller "myapp/controllers/account"
)

func SetupRoutes(app fiber.Router,app_name string) {
  acc := app.Group(app_name)
  acc.Get("/:token",controller.Checker)
  acc.Post("/register",controller.Register)
  acc.Post("/login",controller.Login)
}