package api_routes

import (
"github.com/gofiber/fiber/v2"
controller "myapp/controllers/api"
account "myapp/routes/account"
)

func SetupRoutes(app *fiber.App,app_name string) {
api := app.Group(app_name)
api.Get("/",controller.Home)
account.SetupRoutes(api,"/account")
}