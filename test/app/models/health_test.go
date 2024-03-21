package models_test

import (
	"testing"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app/models"

	"github.com/stretchr/testify/assert"
)

func TestHealthResponseOk(t *testing.T) {
	t.Run("HealthResponseOk", func(t *testing.T) {
		payload := models.HealthResponseOk()

		assert.Equal(t, "ok", payload.Status)
	})
}
