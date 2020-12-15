package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/:p", func(c *fiber.Ctx) {
		p := c.Params("p")
		c.Send(p)
	})

	app.Listen(3000)
}
