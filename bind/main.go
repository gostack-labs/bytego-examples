package main

import (
	"log"

	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego-examples/common"
	"github.com/gostack-labs/bytego/middleware/logger"
)

type School struct {
	Name string `form:"schname"`
}
type City struct {
	CityName string
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
	Age    int    `xml:"age,omitempty" validate:"gte=0,lte=60"`
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
			return common.NewCommonError(1, err.Error())
		}
		return c.JSON(200, s)
	})

	log.Fatal(app.Run(":8000"))
}
