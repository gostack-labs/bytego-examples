package main

import (
	"github.com/gostack-labs/bytego"
)

func main() {
	app := bytego.New()

	app.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, world!")
	})

	_ = app.Run(":8080")
}
