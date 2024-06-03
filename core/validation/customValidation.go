package validation

import (
	"mime/multipart"
	"strings"

	"github.com/asaskevich/govalidator"
)

func InitCustomValidation() {
	govalidator.CustomTypeTagMap.Set("fileImage", govalidator.CustomTypeValidator(isFileImage))
}

func isFileImage(i interface{}, o interface{}) bool {
	// Then, perform a type assertion to extract the *multipart.FileHeader
	file, ok := i.(*multipart.FileHeader)
	if !ok {
		// If the type assertion fails, the value is not a *multipart.FileHeader
		return true
	}

	// Check if the file is not empty
	if file.Size == 0 || file.Filename == "" {
		return true
	}

	// bs, _ := json.Marshal(file)
	// fmt.Println("bs", string(bs))

	// Check the MIME type to ensure it's an image file
	mimeType := file.Header.Get("Content-Type")
	return strings.HasPrefix(mimeType, "image/")
}
