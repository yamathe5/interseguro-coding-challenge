package api

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"interseguro/matrix-api/internal/domain/matrix"
)

type input struct {
	Matrix    json.RawMessage
	Operation string
	Direction string
}

func matrixEndpoint(run func(*fiber.Ctx, input, [][]float64) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req input
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return fail(c, 400, "INVALID_JSON", "Expected a JSON object")
		}
		values, err := matrix.Decode(req.Matrix)
		if err != nil {
			return fail(c, 422, "INVALID_MATRIX", err.Error())
		}
		return run(c, req, values)
	}
}
