package models_test

import (
	"testing"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	t.Run("Create new device model", func(t *testing.T) {
		payload := models.CreateDevice("1", "Device 1", "10.10.20.35")

		assert.Equal(t, "1", payload.Id)
		assert.Equal(t, "Device 1", payload.Name)
		assert.Equal(t, "10.10.20.35", payload.Address)
	})
}
