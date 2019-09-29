package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type Person struct{
	Age int	`form:"age" validate:"required,gt=10"`	// v9默认识别validate
	Name string	`form:"name" validate:"required"`
	//Address string	`form:"address" validate:"required"`
}

var (
	Uni *ut.UniversalTranslator
	Validate *validator.Validate
)

func main(){
	Validate = validator.New()
	zh := zh2.New()
	en := en2.New()
	Uni = ut.New(zh, en)

	r:=gin.Default()
	r.GET("/test", func(context *gin.Context) {
		locale := context.DefaultQuery("locale", "zh")
		trans,_ := Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person
		if err:=context.ShouldBind(&person); err!=nil {
			context.String(500, "%v", err)
			context.Abort()
			return
		}
		if err:=Validate.Struct(person); err!=nil {
			errs:=err.(validator.ValidationErrors)
			sliceErrs:=[]string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			context.String(500, "%v", sliceErrs)
			context.Abort()
			return
		}
		context.String(200, "%v", person)
	})
	r.Run()
}