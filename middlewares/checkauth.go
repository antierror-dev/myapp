package middleware

import (
"github.com/gofiber/fiber/v2"
"myapp/models"
"myapp/database"
"time"
)

func CheckTokenExpire() fiber.Handler {
  return func(c *fiber.Ctx) error {
    token := c.Get("auth")
    if len(token) < 7 {
      response := fiber.Map{
        "status":false,
        "message":"token is not set in headers",
      }
      return c.Status(fiber.StatusConflict).JSON(response)
    }
    
    db := database.GetDB()
    
    var user models.User
    res := db.Where("token = ?", token).First(&user)
    if res.Error != nil {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status":false,"error": res.Error.Error()})
    }
    
    if time.Now().After(user.Expires) {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token is expired","status":false})
    }
    
    c.Locals("user",user)
    return c.Next()
    }
}
