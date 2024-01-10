package routes

import (
  "github.com/gofiber/fiber/v2"
  "myapp/controllers"
  "fmt"
  api "myapp/routes/api"
)

func SetupRoutes(app *fiber.App) {
  app.Get("/", controller.Home)
  api.SetupRoutes(app,"/api")
  fmt.Println("set up routes is finished")
}