package main

import (
	"time"

	"github.com/flangrys/catalyst-portal/docker"
	"github.com/flangrys/catalyst-portal/handlers"
	"github.com/flangrys/catalyst-portal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"
)

const (
	AppVersion = "0.1.0"
	AppAddress = ":8000"
)

func main() {

	docker.Init()

	log.SetLevel(log.LevelInfo)

	app := fiber.New(fiber.Config{
		AppName: "Catalyst Portal",
		Views:   html.New("./views", ".html"),
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(csrf.New(csrf.Config{
		KeyLookup:         "header:X-Csrf-Token",
		CookieName:        "__Host-csrf_",
		CookieSameSite:    "Lax",
		CookieSecure:      true,
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
		Expiration:        1 * time.Hour,
		KeyGenerator:      utils.UUIDv4,
	}))

	app.Static("/", "./public", fiber.Static{
		MaxAge: 86400,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil, "layouts/index")
	})

	dashboard := app.Group("dashboard")
	dashboard.Use(middlewares.AuthMiddleware, middlewares.LoginMiddleware)
	dashboard.Get("", handlers.HandleDashboard)
	dashboard.Get("container/", handlers.AllContainersHandler)
	dashboard.Get("container/:id", handlers.DetailedContainerHandler)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(404).Render("statuses/client_error", fiber.Map{
			"Title":   " 400 + 4 | Catalyst Portal",
			"Status":  404,
			"Message": "What the hell you're searching for?",
		}, "layouts/index")
	})

	log.Fatal(app.Listen(AppAddress))
}
