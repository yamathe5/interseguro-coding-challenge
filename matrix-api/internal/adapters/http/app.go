package api

import (
	"errors"
	"interseguro/matrix-api/internal/adapters/gonum"
	"interseguro/matrix-api/internal/application"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func fail(c *fiber.Ctx, status int, code, message string) error {
	return c.Status(status).JSON(fiber.Map{"error": fiber.Map{
		"code": code, "message": message, "requestId": c.GetRespHeader("X-Request-ID"),
	}})
}

func New(clients ...application.StatisticsClient) *fiber.App {
	app := fiber.New(fiber.Config{BodyLimit: 256 << 10, ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, DisableStartupMessage: true, ErrorHandler: func(c *fiber.Ctx, err error) error {
		var httpErr *fiber.Error
		if errors.As(err, &httpErr) {
			return fail(c, httpErr.Code, "HTTP_ERROR", httpErr.Message)
		}
		return fail(c, 500, "INTERNAL_ERROR", "Unexpected error")
	}})
	app.Use(requestid.New())
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		if err := c.Next(); err != nil {
			if err := app.ErrorHandler(c, err); err != nil {
				return err
			}
		}
		slog.Info("request", "requestId", c.GetRespHeader("X-Request-ID"), "method", c.Method(), "status", c.Response().StatusCode(), "durationMs", time.Since(start).Milliseconds())
		return nil
	})
	health := func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) }
	app.Get("/health/live", health)
	app.Get("/health/ready", func(c *fiber.Ctx) error {
		if len(clients) > 0 && clients[0].Ready(c.UserContext()) != nil {
			return fail(c, 503, "NOT_READY", "Statistics service is unavailable")
		}
		return health(c)
	})
	if len(clients) > 0 {
		app.Post("/api/v1/matrices/analyze", matrixEndpoint(analyze(gonum.QR{}, clients[0])))
	}
	app.Post("/api/v1/matrices/rotate", matrixEndpoint(rotate))
	app.Post("/api/v1/matrices/qr", matrixEndpoint(qr(gonum.QR{})))
	return app
}
