package middleware
import (
"github.com/gofiber/fiber/v2"
"fmt"
)
func Setup(app *fiber.App) {
  app.Use(Logger())
  //app.Use(Execute_Time)
  //app.Use(Limiter())
  //app.Use(Cache())
  //app.Use(Compress())
  fmt.Println("all middlewares was set")
}