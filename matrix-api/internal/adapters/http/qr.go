package api

import (
	"github.com/gofiber/fiber/v2"
	"interseguro/matrix-api/internal/application"
)

func qr(f application.Factorizer) func(*fiber.Ctx, input, [][]float64) error {
	return func(c *fiber.Ctx, _ input, values [][]float64) error {
		result, err := application.FactorizeQR(values, f)
		if err != nil {
			return fail(c, 422, "NUMERIC_OVERFLOW", err.Error())
		}
		return c.JSON(result)
	}
}
