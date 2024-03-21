package operations_test

import (
	"testing"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/operations"
	"github.com/stretchr/testify/assert"
)

func TestGetAllDevices(t *testing.T) {
	t.Run("Get all devices Operation", func(t *testing.T) {
		devices := operations.ListAllDevices()

		assert.Equal(t, 2, len(devices))

		assert.Equal(t, "1", devices[0].Id)
		assert.Equal(t, "Device 1", devices[0].Name)
		assert.Equal(t, "10.10.1.1", devices[0].Address)

		assert.Equal(t, "2", devices[1].Id)
		assert.Equal(t, "Device 2", devices[1].Name)
		assert.Equal(t, "10.10.1.2", devices[1].Address)
	})
}
