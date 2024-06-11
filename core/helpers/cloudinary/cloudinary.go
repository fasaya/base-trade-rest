package cloudinary

import (
	cloudinaryConfig "base-trade-rest/core/config/cloudinary"
	"bytes"
	"context"
	"crypto/rand"
	"math/big"

	"io"
	"mime/multipart"
	"path"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add Cloudinary product environment credentials.
	cld, err := cloudinary.NewFromParams(cloudinaryConfig.EnvCloudName(), cloudinaryConfig.EnvCloudAPIKey(), cloudinaryConfig.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	// Convert file
	fileReader, err := convertFile(fileHeader)
	if err != nil {
		return "", err
	}

	// Upload file
	uploadParam, err := cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: generateRandomStringWithDatetime(),
		Folder:   cloudinaryConfig.EnvCloudUploadFolder(),
	})

	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into an in-memory buffer
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	// Create a bytes.Reader from the buffer
	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil
}

func RemoveExtension(filename string) string {
	return path.Base(filename[:len(filename)-len(path.Ext(filename))])
}

// generateRandomString generates a random string of given length.
func generateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteArr := make([]byte, 15)
	for i := range byteArr {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		byteArr[i] = charset[num.Int64()]
	}
	return string(byteArr)
}

// generateRandomStringWithDatetime generates a random string and appends the current datetime.
func generateRandomStringWithDatetime() string {
	randomStr := generateRandomString()
	datetimeStr := time.Now().Format("20060102150405")
	return randomStr + datetimeStr
}
