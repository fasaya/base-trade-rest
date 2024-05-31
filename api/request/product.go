package request

type ProductCreateRequest struct {
	Name string `json:"name" form:"name"`
	File string `json:"file" form:"file"`
}
