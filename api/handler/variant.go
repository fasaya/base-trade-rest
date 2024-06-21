package handler

import (
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VariantHandler struct {
	VariantService service.IVariantService
	Validate       *validator.Validate
}

func NewVariantHandler(variantService service.IVariantService) VariantHandler {
	var VariantHandler = VariantHandler{
		VariantService: variantService,
		Validate:       helpers.Validate,
	}
	return VariantHandler
}

func (h VariantHandler) Index(ctx *gin.Context) {

}

func (h VariantHandler) Store(ctx *gin.Context) {

}

func (h VariantHandler) Show(ctx *gin.Context) {

}

func (h VariantHandler) Update(ctx *gin.Context) {

}

func (h VariantHandler) Delete(ctx *gin.Context) {

}
