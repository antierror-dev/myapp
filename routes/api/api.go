package api_routes

import (
"github.com/gofiber/fiber/v2"
controller "myapp/controllers/api"
)

func SetupRoutes(app *fiber.App,app_name string) {
api := app.Group(app_name)
api.Get("/",controller.Home)
}