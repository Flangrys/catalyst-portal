package handlers

import (
	"github.com/flangrys/catalyst-portal/docker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleDashboard(c *fiber.Ctx) error {

	containers, err := docker.GetRunningContainers()

	if err != nil {
		log.Errorf("an error ocurred during container fetching: %w:", err)

		c.Status(403).Render("statuses/client_error.html", fiber.Map{
			"Title":   "Hum? | Catalyst Portal",
			"Status":  403,
			"Message": "It seems that this resource was well hid.",
		})
	}

	return c.Render("dashboard/dashboard", fiber.Map{
		"Title":      "All Containers | Catalyst Portal",
		"Count":      len(containers),
		"Containers": containers,
	}, "layouts/index")
}

func AllContainersHandler(c *fiber.Ctx) error {
	return c.Status(200).Render("dashboard/detailed_dashboard", nil, "layouts/index")
}

func DetailedContainerHandler(c *fiber.Ctx) error {

	container_id := c.Params("id")

	containers, err := docker.GetDetailedContainer(container_id)

	if err != nil {
		log.Error(err)

		return c.Status(404).Render("statuses/client_error", fiber.Map{
			"Title": "Hum? | Catalyst Portal",
			"Err":   err,
		}, "layouts/index")
	}

	container := containers[0]

	return c.Render("dashboard/detailed_container", fiber.Map{
		"Title":  "Container | Catalyst Portal",
		"ID":     container.ID[:10],
		"Name":   container.Names[0],
		"Image":  container.Image,
		"State":  container.State,
		"Status": container.Status,
	}, "layouts/index")
}
