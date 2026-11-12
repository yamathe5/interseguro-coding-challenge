package api

import (
	"github.com/gofiber/fiber/v2"
	"interseguro/matrix-api/internal/application"
)

func rotate(c *fiber.Ctx, req input, values [][]float64) error {
	if req.Direction != "" && req.Direction != "clockwise" {
		return fail(c, 422, "INVALID_DIRECTION", "Only clockwise rotation is supported")
	}
	return c.JSON(application.RotateMatrix(values))
}
