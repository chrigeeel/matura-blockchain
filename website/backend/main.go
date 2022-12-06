package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("../frontend/build", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New())

	app.Static("/static", "../frontend/build/static")

	app.Use("/", HandleRender)

	app.Listen(":6969")
}

func HandleRender(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
