package main

import (
	"fmt"
	"log"

	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego/middleware/logger"
	"github.com/gostack-labs/bytego/middleware/recovery"
)

func main() {
	app := bytego.New()
	app.Use(logger.New(), recovery.New())
	app.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, world!"+c.RoutePath())
	})
	g := app.Group("/group", func(c *bytego.Ctx) error {
		fmt.Println("group middleware start")
		fmt.Println(c.RoutePath())
		err := c.Next()
		fmt.Println("group middleware done")
		return err
	})
	g.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, group")
	})
	g.GET("/a", func(c *bytego.Ctx) error {
		return c.String(200, "hello, group a")
	})
	log.Fatal(app.Run(":8000"))
}
