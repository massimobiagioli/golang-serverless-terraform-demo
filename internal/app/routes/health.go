package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/models"
)

func handleHealth(c *fiber.Ctx) error {
	return c.JSON(models.HealthResponseOk())
}

func MountHealthRoutes(router fiber.Router) {
	router.Get("/health", handleHealth)
}
