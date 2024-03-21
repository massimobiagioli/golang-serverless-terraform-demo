package routes_test

import (
	"net/http/httptest"
	"testing"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	fiberApp := app.Build()

	t.Run("Health Check", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/health", nil)
		resp, _ := fiberApp.Test(req, 1)

		assert.Equal(t, 200, resp.StatusCode)
	})
}
