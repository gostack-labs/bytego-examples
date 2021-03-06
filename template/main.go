package main

import (
	"log"

	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego/middleware/logger"
)

func main() {
	app := bytego.New()
	app.Use(logger.New())
	app.SetRender(bytego.NewTemplate("views/*.html"))
	app.GET("/", func(c *bytego.Ctx) error {
		return c.View(200, "hello", "world")
	})
	log.Fatal(app.Run(":8000"))
}
