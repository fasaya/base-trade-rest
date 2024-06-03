package request

import "mime/multipart"

type ProductCreateRequest struct {
	Name  string                `json:"name" form:"name" valid:"required"`
	Image *multipart.FileHeader `json:"image" form:"image" valid:"mime=image/*,maxsize=5242880"`
}
