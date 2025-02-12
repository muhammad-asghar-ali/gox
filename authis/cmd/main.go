package main

import (
	"github.com/gofiber/fiber/v2"

	"authis/internal/config"
	"authis/internal/database"
	"authis/internal/handlers"
	"authis/internal/middlewares"
	"authis/internal/models"
)

func main() {
	cfg := config.GetConfig()

	app := fiber.New()

	database.Connect(cfg)
	models.Migrate(database.GetDB())

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	app.Use(middlewares.Auth)
	app.Get("/user", handlers.User)
	app.Post("/logout", handlers.Logout)
	app.Post("/refresh", handlers.Refresh)

	app.Listen(":6000")
}
