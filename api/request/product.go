package request

import "mime/multipart"

type ProductCreateRequest struct {
	Name  string                `json:"name" form:"name"`
	Image *multipart.FileHeader `json:"image" form:"image"`
}
