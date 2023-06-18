package translate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	evValidation "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
	"sync"
)

var transInstance ut.Translator

var lock = &sync.Mutex{}

func RegisterTranslator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		if err := evValidation.RegisterDefaultTranslations(v, DefaultTranslator()); err != nil {
			return err
		}
	}
	return nil
}

func Translate(err error) map[string]string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		result := make(map[string]string, 0)
		for _, validationError := range validationErrors {
			result[validationError.Field()] = validationError.Translate(DefaultTranslator())
		}
		return result
	}
	return map[string]string{
		"Error": err.Error(),
	}
}

func DefaultTranslator() ut.Translator {
	if transInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if transInstance == nil {
			enLang := en.New()
			uni := ut.New(enLang, enLang)
			transInstance, _ = uni.GetTranslator("en")
		}
	}
	return transInstance
}
