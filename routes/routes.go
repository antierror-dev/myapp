package routes

import (
  "github.com/gofiber/fiber/v2"
  "myapp/controllers"
  "fmt"
)

func SetupRoutes(app *fiber.App) {
  app.Get("/", controller.Home)
  fmt.Println("set up routes is finished")
}