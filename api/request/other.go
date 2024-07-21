package request

type FieldValueRequest struct {
	Field string
	Value any
}

type PaginateRequest struct {
	Page   int    `json:"page" form:"page" validate:"number"`
	Search string `json:"search" form:"search"`
	Limit  int    `json:"limit" form:"limit" validate:"number"`
	// Sort   string `json:"sort" form:"sort"`
}
