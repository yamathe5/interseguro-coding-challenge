package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"interseguro/matrix-api/internal/application"
)

func analyze(f application.Factorizer, client application.StatisticsClient) func(*fiber.Ctx, input, [][]float64) error {
	return func(c *fiber.Ctx, req input, values [][]float64) error {
		result, err := application.AnalyzeMatrix(c.UserContext(), values, req.Operation, c.GetRespHeader("X-Request-ID"), f, client)
		switch {
		case errors.Is(err, application.ErrDependency):
			return fail(c, 502, "DEPENDENCY_FAILURE", err.Error())
		case errors.Is(err, application.ErrOperation):
			return fail(c, 422, "INVALID_OPERATION", err.Error())
		case err != nil:
			return fail(c, 422, "NUMERIC_OVERFLOW", err.Error())
		default:
			return c.JSON(result)
		}
	}
}
