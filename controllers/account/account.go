package account_controller

import (
  "github.com/gofiber/fiber/v2"
  "myapp/models"
  "myapp/database"
  "time"
  "fmt"
  "errors"
)

func Checker(c *fiber.Ctx) error {
  token := c.Params("token")
  if token == "" || len(token) < 10 {
    return c.Status(fiber.StatusBadRequest).SendString("Error Token")
  }
  var user models.User
  db := database.GetDB()
  db.Where("token = ?",token).First(&user)
  if user.ID == 0{
    return fiber.ErrUnauthorized
  }
  return c.JSON(fiber.Map{"email":user.Email,"expires":user.Expires,})
}

func validation(e string , p string) error{
  if e == ""{
    return errors.New("email is required")
  }
  if p == ""{
    return errors.New("password is required")
  }
  return nil
}

func Register(c *fiber.Ctx) error {
  db := database.GetDB()
  var data map[string]string
  if err := c.BodyParser(&data); err != nil {return err}
  
  if err := validation(data["email"], data["password"]); err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
  }
  
  var user models.User
  db.Where("email = ?",data["email"]).First(&user)
  if user.ID != 0{
    return c.Status(fiber.StatusConflict).SendString("Email already in use")
  }
  user = models.User{
    Email: data["email"],
    Password: data["password"],
    Expires: time.Now(),
  }
  result := db.Create(&user)
  if result.Error != nil {
    return c.Status(fiber.StatusConflict).SendString("Error on creating user")
  }
  return c.JSON(user)
}



const (
  tokenExpirationDuration = 5 * 24 * time.Hour
)

func Login(c *fiber.Ctx) error {
  db := database.GetDB()
  request := new(struct {
  Email string
  Password string
  })
  if err := c.BodyParser(request); err != nil {
    return err
  }
  user := new(models.User)
  result := db.Where("email = ? AND password = ?", request.Email, request.Password).First(&user)
  if result.Error != nil {
    return fiber.ErrUnauthorized
  }

  user.Token = fmt.Sprintf("%x", time.Now().UnixNano())
  user.Expires = time.Now().Add(tokenExpirationDuration)

  db.Save(&user)

  return c.JSON(user)
}
