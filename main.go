package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

const (
	AppVersion = "0.1.0"
	AppAddress = ":9080"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Catalyst Portal",
		ServerHeader: "X-Embeding-Protocol: enabled",
		Views:        html.New("./templates", ".html"),
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/", "./public", fiber.Static{
		MaxAge: 86400,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index.view", fiber.Map{"Title": "Catalyst Portal | Index"}, "index.layout")
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Render("404.view", fiber.Map{}, "index.layout")
	})

	log.Fatal(app.Listen(AppAddress))
}
