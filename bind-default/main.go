package main

import (
	"log"

	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego-examples/common/xresult"
	"github.com/gostack-labs/bytego/middleware/logger"
)

func main() {
	app := bytego.New()
	app.Use(logger.New())
	app.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, world!")
	})

	type DefaultA struct {
		DemoA string `default:"demo1"`
	}
	type DefaultB struct {
		DemoB string `default:"demo2"`
	}
	type DefaultC struct {
		DemoC string `default:"demo3"`
	}
	type DefaultD struct {
		DemoD string `default:"demo4"`
	}
	type Default struct {
		Int      int    `form:"int" default:"5"`
		String   string `form:"string" default:"string1"`
		UInt     uint   `form:"uint" default:"11"`
		Int32    int32  `default:"32"`
		Int64    int64  `default:"64"`
		Inta     *int   `default:"65"`
		Bool     bool   `default:"true"`
		DefaultA DefaultA
		DefaultB
		DefaultC *DefaultC `default:"new"`
		*DefaultD
		Slice       []int    `default:"1,2,3"`
		StringSlice []string `default:"a,b,c"`
	}

	// curl -d 'int=10'  http://localhost:8000/bind/default/1
	app.POST("/bind/default/:id", func(c *bytego.Ctx) error {
		var s Default
		if err := c.Bind(&s); err != nil {
			return xresult.Fail(400, err.Error())
		}
		return c.JSON(200, xresult.Success(s))
	})

	log.Fatal(app.Run(":8000"))
}
