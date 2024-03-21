package operations

import (
	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/models"
)

func ListAllDevices() []models.Device {
	return []models.Device{
		models.CreateDevice("1", "Device 1", "10.10.1.1"),
		models.CreateDevice("2", "Device 2", "10.10.1.2"),
	}
}
