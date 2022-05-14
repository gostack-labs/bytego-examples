package main

import (
	"log"

	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego-examples/common/xresult"
	"github.com/gostack-labs/bytego/middleware/logger"
)

type School struct {
	Name string `form:"schname"`
}
type City struct {
	CityName string `form:"cityname"`
}
type WantJob struct {
	JobName string
}
type People struct {
	Name   string  `json:"name,omitempty"`
	Parent *People `json:"parent,omitempty"`
}
type Student struct {
	Name   string `xml:"name,omitempty" form:"formname"`
	Age    int    `xml:"age,omitempty" form:"age" validate:"gte=0,lte=60"`
	School School `form:"sch"`
	City
	*WantJob
	Parent     *People
	Header1    string   `header:"request-id"`
	Query1     string   `query:"query1"`
	Param1     string   `param:"id"`
	LikeColors []string `form:"colors[]"`
}

func main() {
	app := bytego.New()
	app.Use(logger.New())
	app.GET("/", func(c *bytego.Ctx) error {
		return c.String(200, "hello, world!")
	})

	//curl -d '{"name":"a","age":22}' -H 'content-type:application/json' http://localhost:8000/bind/student/1
	//curl -d '<student><name>test</name><age>18</age></student>' -H 'content-type:application/xml' http://localhost:8000/bind/student/1
	//curl -d 'formname=test&age=18&sch.schname=aa' -H 'content-type:application/x-www-form-urlencoded' http://localhost:8000/bind/student/1
	//curl -d 'formname=test&age=18&sch.schname=aa'  http://localhost:8000/bind/student/1
	//curl -d 'formname=test&age=18&sch.schname=aa&cityname=hz&jobname=programer&parent.name=pname&parent.parent.name=ppname&colors[]=1&colors[]=2'  http://localhost:8000/bind/student/1
	app.POST("/bind/student/:id", func(c *bytego.Ctx) error {
		var s Student
		if err := c.Bind(&s); err != nil {
			return xresult.Fail(400, err.Error())
		}
		return c.JSON(200, xresult.Success(s))
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
