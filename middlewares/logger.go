package middleware

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() fiber.Handler {
  return logger.New(logger.Config{
    Format:"${time}-[${ip}]- ${method}-${status}: ${path} |${latency} | (${pid})\n",
    TimeZone:   "Asia/Tehran",
  })
}