package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"

	"golang.org/x/exp/slices"

	"strconv"
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
	Validate     *validator.Validate
	Translator   ut.Translator
	maxImageSize = 5 // In MB
)

func init() {
	fmt.Println("Initializing Translator...")

	// Initializing the translator
	enTranslator := en.New()
	uni := ut.New(enTranslator, enTranslator)

	var found bool
	Translator, found = uni.GetTranslator("en")
	if !found {
		panic("Translator not found")
	}

	// Initializing the validator
	Validate = validator.New()

	// Registering the default translations
	en_translations.RegisterDefaultTranslations(Validate, Translator)

	// Registering the custom rule
	err := Validate.RegisterValidation("maxImageSizeInMb", maxImageSizeInMbValidator)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	err = Validate.RegisterValidation("commonImageType", isCommonImageType)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	// Register the custom error message for the "maxImageSizeInMb" validation rule
	Validate.RegisterTranslation("maxImageSizeInMb", Translator, func(ut ut.Translator) error {
		return ut.Add("maxImageSizeInMb", "{0} must not exceed {1} bytes", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("maxImageSizeInMb", fe.Field(), fe.Param())
		return t
	})

	// Register the custom error message for the "commonImageType" validation rule
	Validate.RegisterTranslation("commonImageType", Translator, func(ut ut.Translator) error {
		return ut.Add("commonImageType", "{0} must be a jpg, jpeg or png filetype", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("commonImageType", fe.Field(), fe.Param())
		return t
	})
}

func getValidationErrors(err error, trans *ut.Translator) map[string]string {
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

// TO DO: Need to fix, it still doesn't work
func ValidateImage(v *validator.Validate, fileHeader *multipart.FileHeader) error {
	if err := v.Var(fileHeader, "commonImageType"); err != nil {
		return fmt.Errorf("file format must be JPG or JPEG")
	}

	if err := v.Var(fileHeader, "maxImageSizeInMb=5"); err != nil {
		return fmt.Errorf("image size cannot exceed 2MB")
	}
	return nil
}

func ValidateImageUpload(file *multipart.FileHeader) error {
	// Validating if the current field's value contains the path to a valid image file
	allowedExtensions := []string{"jpg", "jpeg", "png", "svg"}
	ext := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])
	if !slices.Contains(allowedExtensions, ext) {
		return errors.New("file format must be JPG, JPEG or PNG")
	}

	// Validating the maximum size of an image
	if file.Size <= int64(maxImageSize) {
		return errors.New("image size cannot exceed 5MB")
	}

	return nil
}

// maxImageSizeInMbValidator is a custom validator for validating the maximum size of an image
func maxImageSizeInMbValidator(fl validator.FieldLevel) bool {
	file, ok := fl.Top().Interface().(*multipart.FileHeader)
	if !ok {
		return false
	}

	// maxSize := int64(2) // 2MB

	maxSize := fl.Param() // Get the maximum size parameter from the validation tag
	size, err := strconv.ParseInt(maxSize, 10, 32)
	if err != nil {
		// Handle the case where the maximum size parameter is not a valid integer
		return false
	}

	// Convert the maximum size from MB to bytes
	size = size * 1024 * 1024

	return file.Size <= size

}

// isCommonImageType is the validation function for validating if the current field's value contains the path to a valid image file
func isCommonImageType(fl validator.FieldLevel) bool {
	file, ok := fl.Top().Interface().(*multipart.FileHeader)
	if !ok {
		return false
	}

	ext := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])

	return ext == "jpg" || ext == "jpeg" || ext == "png"
}
