package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/routes"
)

func Build() *fiber.App {
	fiberApp := fiber.New()

	apiGroup := fiberApp.Group("/api")

	routes.MountHealthRoutes(apiGroup)
	routes.MountDeviceRoutes(apiGroup)

	return fiberApp
}
