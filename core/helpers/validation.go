package helpers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// var Validate *validator.Validate
// var Translator ut.Translator

var (
	Validate   *validator.Validate
	Translator ut.Translator
)

func init() {
	fmt.Println("Initializing Translator...")

	enTranslator := en.New()
	uni := ut.New(enTranslator, enTranslator)

	var found bool
	Translator, found = uni.GetTranslator("en")
	if !found {
		panic("Translator not found")
	}

	Validate = validator.New()

	en_translations.RegisterDefaultTranslations(Validate, Translator)
}

func getValidationErrors(err error, trans *ut.Translator) map[string]string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return nil
	}

	var errors = make(map[string]string)

	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errors[strings.ToLower(e.Field())] = e.Translate(*trans)
	}

	return errors
}

func HandleValidationError(ctx *gin.Context, err error, trans *ut.Translator) {
	CreateValidationErrorResponse(ctx, getValidationErrors(err, trans))
}
