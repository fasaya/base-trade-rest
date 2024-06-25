package cloudinary

import (
	"os"
)

func EnvCloudName() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file on EnvCloudName")
	// }
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file on EnvCloudAPIKey")
	// }
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file on EnvCloudAPISecret")
	// }
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file EnvCloudUploadFolder")
	// }
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
