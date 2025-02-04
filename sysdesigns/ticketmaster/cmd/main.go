package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/routes"
)

func main() {
	app := fiber.New()
	port := config.GetConfig().GetServerPort()

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(port))
}
