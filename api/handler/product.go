package handler

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"
	"base-trade-rest/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	ProductService service.IProductService
	Validate       *validator.Validate
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	var productHandler = ProductHandler{
		ProductService: productService,
		Validate:       helpers.Validate,
	}
	return &productHandler
}

func (h *ProductHandler) Index(ctx *gin.Context) {
	products, err := h.ProductService.GetListProduct()

	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Products successfully fetched", products)
}

func (h *ProductHandler) Store(ctx *gin.Context) {
	var request request.ProductCreateRequest

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

	if request.Image.Size != 0 || request.Image.Filename != "" {

		// TO DO: Need to fix, it still doesn't work
		err = helpers.ValidateImage(h.Validate, request.Image)
		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		// Extract the filename without extension
		// fileName := helpers.RemoveExtension(request.Image.Filename)

		// uploadResult, err := helpers.UploadFile(request.Image, fileName)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
	}

	// Get authenticated user
	userData := helpers.GetAuthUser(ctx)

	product := model.Product{
		Name:     request.Name,
		ImageURL: nil,
		UserID:   userData.ID,
	}

	newUser, err := h.ProductService.CreateProduct(&product)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Product successfully created", newUser)
}

func (h *ProductHandler) Show(ctx *gin.Context) {

}

func (h *ProductHandler) Update(ctx *gin.Context) {

}

func (h *ProductHandler) Delete(ctx *gin.Context) {

}
