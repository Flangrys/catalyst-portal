package handlers

import (
	"github.com/flangrys/catalyst-portal/docker"
	"github.com/gofiber/fiber/v2"
)

func DashboardHandler(c *fiber.Ctx) error {

	dummy_containers, containers_count := docker.GetRunningContainers()

	return c.Render("dashboard/dashboard", fiber.Map{
		"ContainersCount": containers_count,
		"ContainersDummy": dummy_containers,
	}, "layouts/index")
}
