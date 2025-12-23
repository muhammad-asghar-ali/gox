package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/routes"
)

func main() {
	app := fiber.New()
	port := config.GetConfig().GetServerPort()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New())
	app.Use(cache.New())

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(port))
}
