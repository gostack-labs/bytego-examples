package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego-examples/common/xresult"
	"github.com/gostack-labs/bytego/middleware/logger"
)

type validateError struct {
	err validator.FieldError
}

func (q validateError) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("'%s' is %s", q.err.Field(), q.err.ActualTag()))
	// condition parameters
	if q.err.Param() != "" {
		sb.WriteString(" " + q.err.Param())
	}
	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", but got %v", q.err.Value()))
	}
	return sb.String()
}

func Translate(ctx *bytego.Ctx, err error) error {
	var trans ut.Translator
	var isEN bool
	lang := ctx.Header("Accept-Language")
	if lang != "" {
		lang = strings.Split(lang, ",")[0]
	}
	switch lang {
	case "zh-CN":
		trans, _ = uni.GetTranslator("zh")
	case "en-US":
		trans, _ = uni.GetTranslator("en")
		isEN = true
	default:
		trans, _ = uni.GetTranslator("zh")
	}
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var ret strings.Builder
		for _, err := range validationErrors {
			if isEN {
				ret.WriteString(validateError{err: err}.String())
			} else {
				ret.WriteString(err.Translate(trans))
			}
			ret.WriteString(";")
		}
		return xresult.Fail(10001, ret.String())
	}
	return err
}

type Student struct {
	Name string `form:"name" validate:"required" label:"姓名" json:"name,omitempty"`
	Age  int    `form:"age,omitempty" validate:"gte=18,lte=30" label:"年龄" json:"age,omitempty"`
}

var uni *ut.UniversalTranslator

func main() {
	app := bytego.New()
	app.Debug(true)
	app.Use(logger.New())

	uni = ut.New(en.New(), zh.New())
	validator := validator.New()
	trans, _ := uni.GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(validator, trans)
	//register a tag as filed name
	// validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
	// 	name := fld.Tag.Get("label")
	// 	return name
	// })
	app.SetValidator(validator.Struct, Translate)

	// simple validator
	// app.Validator(validator.New().Struct)

	//curl -d 'name=&age=60'  http://localhost:8000/
	// curl  -H 'Accept-Language: zh-CN' -d 'name=&age=60'  http://localhost:8000/
	// curl  -H 'Accept-Language: en-US' -d 'name=&age=60'  http://localhost:8000/
	app.POST("/", func(c *bytego.Ctx) error {
		var s Student
		if err := c.Bind(&s); err != nil {
			return err
		}
		return c.JSON(200, s)
	})

	log.Fatal(app.Run(":8000"))
}
