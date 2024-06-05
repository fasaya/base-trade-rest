package request

import "mime/multipart"

type ProductCreateRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
	// Image *multipart.FileHeader `json:"image" form:"image" validate:"commonImageType,maxImageSizeInMb=5"`
	Image *multipart.FileHeader `json:"image" form:"image"`
}
