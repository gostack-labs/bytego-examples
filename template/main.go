package main

import (
	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego/middleware/logger"
)

func main() {
	app := bytego.New()
	app.Use(logger.New())
	app.Render(bytego.NewDefaultTemplate("views/*.html"))

	app.GET("/", func(c *bytego.Ctx) error {
		return c.View(200, "hello", "world")
	})

	_ = app.Run(":8080")
}
