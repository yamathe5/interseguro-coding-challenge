package api

import (
	"errors"
	"interseguro/matrix-api/internal/adapters/gonum"
	"interseguro/matrix-api/internal/application"
	"log/slog"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var requestIDPattern = regexp.MustCompile("^[A-Za-z0-9_-]{1,64}$")

func fail(c *fiber.Ctx, status int, code, message string) error {
	return c.Status(status).JSON(fiber.Map{"error": fiber.Map{
		"code": code, "message": message, "requestId": c.GetRespHeader("X-Request-ID"),
	}})
}

func New(client application.StatisticsClient) *fiber.App {
	app := fiber.New(fiber.Config{BodyLimit: 256 << 10, ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, DisableStartupMessage: true, ErrorHandler: func(c *fiber.Ctx, err error) error {
		var httpErr *fiber.Error
		if errors.As(err, &httpErr) {
			return fail(c, httpErr.Code, "HTTP_ERROR", httpErr.Message)
		}
		return fail(c, 500, "INTERNAL_ERROR", "Unexpected error")
	}})
	app.Use(func(c *fiber.Ctx) error {
		if !requestIDPattern.MatchString(c.Get("X-Request-ID")) {
			c.Request().Header.Del("X-Request-ID")
		}
		return c.Next()
	})
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
	app.Use(recover.New())
	health := func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) }
	app.Get("/health/live", health)
	app.Get("/health/ready", func(c *fiber.Ctx) error {
		if client != nil && client.Ready(c.UserContext()) != nil {
			return fail(c, 503, "NOT_READY", "Statistics service is unavailable")
		}
		return health(c)
	})
	if client != nil {
		app.Post("/api/v1/matrices/analyze", matrixEndpoint(analyze(gonum.QR{}, client)))
	}
	app.Post("/api/v1/matrices/rotate", matrixEndpoint(rotate))
	app.Post("/api/v1/matrices/qr", matrixEndpoint(qr(gonum.QR{})))
	return app
}
