package user_controller

import (
"github.com/gofiber/fiber/v2"
"myapp/models"
"myapp/database"
)

func ChangeNames(c *fiber.Ctx) error {
  var data map[string]string
  if err := c.BodyParser(&data); err != nil {return err}
  var name,family string
  name = data["name"]
  family = data["family"]
  if len(name) < 2{
    response := fiber.Map{
      "status":false,
      "message":"please send a valid name",
    }
    return c.JSON(response)
  }
  db := database.GetDB()
  var user models.User
  user,_ = c.Locals("user").(models.User)
  user.Fname = name
  user.Lname = family
  db.Save(&user)
  response := fiber.Map{
    "status":true,
    "message":"new and family is set",
    "user":user,
  }
  return c.JSON(response)
}

