package api

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
)

//go:embed openapi.json
var openAPISpec []byte

//go:embed swagger.html
var swaggerPage []byte

func registerDocs(app *fiber.App) {
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Type("html", "utf-8").Send(swaggerPage)
	})
	app.Get("/openapi.json", func(c *fiber.Ctx) error {
		return c.Type("json").Send(openAPISpec)
	})
}
