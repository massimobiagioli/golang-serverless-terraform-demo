package routes_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app"
	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/models"
)

func TestGetAllDevices(t *testing.T) {
	fiberApp := app.Build()

	t.Run("Get all devices", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/device", nil)
		resp, _ := fiberApp.Test(req, 1)

		assert.Equal(t, 200, resp.StatusCode)

		var devices []models.Device
		err := json.NewDecoder(resp.Body).Decode(&devices)
		assert.NoError(t, err)
		assert.NotEmpty(t, devices)

		assert.Equal(t, "1", devices[0].Id)
		assert.Equal(t, "Device 1", devices[0].Name)
		assert.Equal(t, "10.10.1.1", devices[0].Address)

		assert.Equal(t, "2", devices[1].Id)
		assert.Equal(t, "Device 2", devices[1].Name)
		assert.Equal(t, "10.10.1.2", devices[1].Address)
	})
}
