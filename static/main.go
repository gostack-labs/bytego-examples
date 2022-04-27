package main

import (
	"log"

	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego/middleware/logger"
)

func main() {
	app := bytego.New()
	app.Use(logger.New())
	app.Static("/public", "./public/")
	app.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, world!")
	})
	g := app.Group("/group")
	g.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, group")
	})
	g.Static("/public", "./public/")
	log.Fatal(app.Run(":8000"))
}
