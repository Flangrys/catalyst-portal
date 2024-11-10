package main

import (
	"log"

	"github.com/flangrys/catalyst-portal/docker"
	"github.com/flangrys/catalyst-portal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

const (
	AppVersion = "0.1.0"
	AppAddress = ":9080"
)

func main() {

	docker.Init()

	app := fiber.New(fiber.Config{
		AppName: "Catalyst Portal",
		Views:   html.New("./views", ".html"),
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/", "./public", fiber.Static{
		MaxAge: 86400,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/index")
	})

	app.Get("/dashboard", handlers.DashboardHandler)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Render("404", fiber.Map{}, "layouts/index")
	})

	log.Fatal(app.Listen(AppAddress))
}
