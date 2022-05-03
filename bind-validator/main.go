package main

import (
	"log"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego-examples/common/xresult"
	"github.com/gostack-labs/bytego/middleware/logger"
)

func Translate(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return xresult.Fail(10001, err.Error())
	} else {
		var ret string
		var errCount int = len(validationErrors)
		if errCount == 1 {
			ret += validationErrors[0].Translate(trans) + "。"
		} else {
			for _, e := range validationErrors {
				ret += e.Translate(trans) + ";"
			}
		}
		return xresult.Fail(10002, ret)
	}
}

var trans ut.Translator

type Student struct {
	Name string `form:"name" validate:"required" label:"姓名" json:"name,omitempty"`
	Age  int    `form:"age,omitempty" validate:"gte=18,lte=30" label:"年龄" json:"age,omitempty"`
}

func main() {
	app := bytego.New()
	app.Debug(true)
	app.Use(logger.New())

	// validator with translation
	validator := validator.New()
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	_ = translation.RegisterDefaultTranslations(validator, trans)
	//register a tag as filed name
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	app.SetValidator(validator.Struct, Translate)

	// simple validator
	// app.Validator(validator.New().Struct)

	//curl -d 'name=&age=60'  http://localhost:8000/
	app.POST("/", func(c *bytego.Ctx) error {
		var s Student
		if err := c.Bind(&s); err != nil {
			return err
		}
		return c.JSON(200, s)
	})

	log.Fatal(app.Run(":8000"))
}
