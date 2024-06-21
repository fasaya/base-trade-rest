package request

type VariantCreateRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Quantity  int    `json:"quantity" form:"quantity" validate:"required,number"`
	ProductID string `json:"product_id" form:"product_id" validate:"required"`
}
