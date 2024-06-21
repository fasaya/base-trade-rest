package handler

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"
	"base-trade-rest/core/repository"
	"base-trade-rest/core/service"
	"base-trade-rest/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VariantHandler struct {
	VariantService    service.IVariantService
	Validate          *validator.Validate
	ProductRepository repository.IProductRepository
}

func NewVariantHandler(variantService service.IVariantService) VariantHandler {
	var VariantHandler = VariantHandler{
		VariantService:    variantService,
		Validate:          helpers.Validate,
		ProductRepository: repository.NewProductRepository(database.GetDB()),
	}
	return VariantHandler
}

func (h VariantHandler) Index(ctx *gin.Context) {
	variant, err := h.VariantService.GetListVariant()

	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Data successfully fetched", variant)
}

func (h VariantHandler) Store(ctx *gin.Context) {
	var request request.VariantCreateRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Validate.Struct(request)
	if err != nil {
		helpers.HandleValidationError(ctx, err, &helpers.Translator)
		return
	}

	// Check if product exists
	product, err := h.ProductRepository.GetProductByKey("uuid", request.ProductID)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, "Product not found")
		return
	}

	variant := model.Variant{
		Name:      request.Name,
		Quantity:  request.Quantity,
		ProductID: product.ID,
	}

	newUser, err := h.VariantService.CreateVariant(&variant)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Data successfully created", newUser)
}

func (h VariantHandler) Show(ctx *gin.Context) {

}

func (h VariantHandler) Update(ctx *gin.Context) {

}

func (h VariantHandler) Delete(ctx *gin.Context) {

}
