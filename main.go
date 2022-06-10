package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// HeadersMiddleware headers setup and defination
func HeadersMiddleware(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	c.Set("Content-Type", "application/json")
	return c.Next()

}

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(HeadersMiddleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${pid} ${locals:requestid} ${status} ${latency} - ${method} ${path} ${query:}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "UTC",
	}))

	app.Get("/dashboard", monitor.New())

	for path, handlers := range Routes {
		for method, handler := range handlers {
			switch method {
			case "GET":
				app.Get(path, handler)
			case "POST":
				app.Post(path, handler)
			case "PUT":
				app.Put(path, handler)
			case "PATCH":
				app.Patch(path, handler)
			case "DELETE":
				app.Delete(path, handler)
			}
		}
	}
	return app
}

func main() {
	app := SetupRouter()
	app.Listen(":8080")
}
