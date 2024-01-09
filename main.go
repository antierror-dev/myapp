package main

import (
"github.com/gofiber/fiber/v2"
"myapp/routes"
"myapp/configs"
"myapp/middlewares"
"fmt"
"log"
)

func main(){
app := fiber.New(configs.Appconfig)
middleware.Setup(app)
routes.SetupRoutes(app)
fmt.Println("app is ready to run...")
log.Fatal(app.Listen(":3000"))
}