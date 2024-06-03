package validation

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

func InitCustomValidation() {
	govalidator.CustomTypeTagMap.Set("fileImage", govalidator.CustomTypeValidator(isFileImage))

	govalidator.InterfaceParamTagMap["fileImageMaxSize"] = govalidator.InterfaceParamValidator(isFileImageMaxSize)
	govalidator.InterfaceParamTagRegexMap["fileImageMaxSize"] = regexp.MustCompile(`^fileImageMaxSize\((.*)\)$`)
}

func isFileImageMaxSize(i interface{}, params ...string) bool {
	// Then, perform a type assertion to extract the *multipart.FileHeader
	file, ok := i.(*multipart.FileHeader)
	if !ok {
		// If the type assertion fails, the value is not a *multipart.FileHeader
		return true
	}

	bs, _ := json.Marshal(file)
	fmt.Println("bs", string(bs))

	// Check if the file is not empty
	if file.Size == 0 || file.Filename == "" {
		return true
	}
	fmt.Println("a1")

	// Convert params to int64 to get the maxFileSize
	maxFileSize, err := strconv.ParseInt(params[0], 10, 64)
	if err != nil {
		return false
	}
	fmt.Println("a2")

	// Check if file size is under or equal to maxFileSize
	maxFileSize = maxFileSize * 1024 * 1024 // Default to 5 MB
	if file.Size > maxFileSize {
		return false
	}
	fmt.Println("a3")

	// Check the MIME type to ensure it's an image file
	mimeType := file.Header.Get("Content-Type")
	return strings.HasPrefix(mimeType, "image/")
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

	// Check the MIME type to ensure it's an image file
	mimeType := file.Header.Get("Content-Type")
	return strings.HasPrefix(mimeType, "image/")
}
