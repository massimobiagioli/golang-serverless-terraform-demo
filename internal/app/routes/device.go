package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/operations"
)

func handleListDevices(c *fiber.Ctx) error {
	devices := operations.ListAllDevices()
	return c.JSON(devices)
}

func MountDeviceRoutes(router fiber.Router) {
	deviceGroup := router.Group("/device")

	deviceGroup.Get("/", handleListDevices)
}
