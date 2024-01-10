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

func ChangePhone(c *fiber.Ctx) error {
  var data map[string]string
  if err := c.BodyParser(&data); err != nil {return err}
  var phone string = data["phone"]
  if len(phone) < 10{
    response := fiber.Map{
      "status":false,
      "message":"please send a phone number",
    }
    return c.JSON(response)
  }
  db := database.GetDB()
  var user models.User
  user,_ = c.Locals("user").(models.User)
  user.Phone = phone
  db.Save(&user)
  response := fiber.Map{
    "status":true,
    "message":"new phone number is set",
    "user":user,
  }
  return c.JSON(response)
}


func ChangePassword(c *fiber.Ctx) error{
  var data map[string]string
  if err := c.BodyParser(&data); err != nil {return err}
  var password string = data["password"]
  if len(password) < 6{
    response := fiber.Map{
      "status":false,
      "message":"please send a password",
    }
    return c.JSON(response)
  }
  db := database.GetDB()
  var user models.User
  user,_ = c.Locals("user").(models.User)
  user.Password = password
  db.Save(&user)
  response := fiber.Map{
    "status":true,
    "message":"new password is set",
    "user":user,
  }
  return c.JSON(response)
}


func ChangeMail(c *fiber.Ctx) error {
  var data map[string]string
  if err := c.BodyParser(&data); err != nil {return err}
  var email string = data["email"]
  if len(email) < 5{
    response := fiber.Map{
      "status":false,
      "message":"please send a email",
    }
    return c.JSON(response)
  }
  
  db := database.GetDB()
  var finduser models.User
  db.Where("email = ?",email).First(&finduser)
  if finduser.ID != 0{
  response := fiber.Map{
    "status":false,
    "message":"email already use",
    }
  return c.Status(fiber.StatusConflict).JSON(response)
  }
  
  var user models.User
  user,_ = c.Locals("user").(models.User)
  user.Email = email
  db.Save(&user)
  response := fiber.Map{
    "status":true,
    "message":"new email is set",
    "user":user,
  }
  return c.JSON(response)
}