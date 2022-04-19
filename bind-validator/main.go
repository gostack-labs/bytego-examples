package main

import (
	"log"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego-examples/common"
	"github.com/gostack-labs/bytego/middleware/logger"
)

func Translate(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return common.NewCommonError(10001, err.Error())
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
		return common.NewCommonError(10002, ret)
	}
}

var trans ut.Translator

type Student struct {
	Name string `form:"name" validate:"required" json:"name,omitempty"`
	Age  int    `form:"age,omitempty" validate:"gte=18,lte=30" json:"age,omitempty"`
}

func main() {
	app := bytego.New()
	app.Use(logger.New())

	// validator with translation
	validator1 := validator.New()
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	_ = translation.RegisterDefaultTranslations(validator1, trans)
	app.Validator(validator1.Struct, Translate)

	// simple validator
	// app.Validator(validator.New().Struct)

	//curl -d 'name=&age=60'  http://localhost:8080/
	app.POST("/", func(c *bytego.Ctx) error {
		var s Student
		if err := c.Bind(&s); err != nil {
			return err
		}
		return c.JSON(200, s)
	})

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
