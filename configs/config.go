package configs

import (
  "github.com/gofiber/fiber/v2"
  "github.com/goccy/go-json"
  "fmt"
  //"github.com/gofiber/template/html/v2"
  )



var(
  APP_NAME string = "MyApp"
  SERVER_HEADER string  = "Petros"
  MULTI_CPU bool = true
)

var Appconfig fiber.Config
func init(){
//engine := html.New("./views", ".html")
Appconfig = fiber.Config{
    AppName: APP_NAME,
    ServerHeader: SERVER_HEADER,
    Prefork: MULTI_CPU,
    CaseSensitive: false,
    Concurrency: 256 * 1024,//connected devices len
    DisableStartupMessage: false,
    EnablePrintRoutes: false,
    ReadBufferSize: 4096,
    WriteBufferSize:4096,
    StrictRouting: false,
    //Views: engine,
    JSONEncoder: json.Marshal,
    JSONDecoder: json.Unmarshal,
  }
  fmt.Println("config of app is finished")
}