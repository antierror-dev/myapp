package user_routes

import (
"github.com/gofiber/fiber/v2"
controller "myapp/controllers/user"
middleware "myapp/middlewares"
)

func SetupRoutes(app fiber.Router,app_name string) {
us := app.Group(app_name,middleware.CheckTokenExpire())
us.Get("/",controller.Home)
us.Post("/change-mail",controller.ChangeMail)
us.Post("/change-password",controller.ChangePassword)
us.Post("/change-phone",controller.ChangePhone)
us.Post("/change-names",controller.ChangeNames)
}