package configs

import (
  "github.com/gofiber/fiber/v2"
  "github.com/goccy/go-json"
  "github.com/joho/godotenv"
  "os"
  "fmt"
  "strconv"
  //"github.com/gofiber/template/html/v2"
)


func parsebool(str string) (bool) {
  res,_ := strconv.ParseBool(str)
  return res
}


var (
  APP_NAME string
  SERVER_HEADER string
  MULTI_CORE bool
  Appconfig fiber.Config
)

func loadenv(){
  err := godotenv.Load(".env")
  if err!=nil{
    panic("set .env file")
  }
  APP_NAME = os.Getenv("APP_NAME")
  SERVER_HEADER = os.Getenv("SERVER_HEADER")
  MULTI_CORE = parsebool(os.Getenv("MULTI_CORE"))
}


func init(){
  loadenv()
  Appconfig = fiber.Config{
    AppName: APP_NAME,
    ServerHeader: SERVER_HEADER,
    Prefork: MULTI_CORE,
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
//engine := html.New("./views", ".html")