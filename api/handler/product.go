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
		// err = helpers.ValidateImage(h.Validate, request.Image) // TO DO: Need to fix, it still doesn't work
		err = helpers.ValidateImageUpload(request.Image)
		if err != nil {
			helpers.CreateValidationErrorResponse(ctx, gin.H{"image": err.Error()})
			return
		}

		// Upload Image
		// fileName := helpers.RemoveExtension(request.Image.Filename)
		// ...
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
	productUUID := ctx.Param("productUUID")

	product, err := h.ProductService.GetDetailProductByUUID(productUUID)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, "Data not found")
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Product successfully fetched", product)
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	productUUID := ctx.Param("productUUID")
	var request request.ProductUpdateRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.ProductService.GetDetailProductByUUID(productUUID)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	updatedProduct, err := h.ProductService.UpdateProduct(product)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Product successfully updated", updatedProduct)
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	productUUID := ctx.Param("productUUID")

	// Delete product
	err := h.ProductService.DeleteProductByUUID(productUUID)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Product successfully deleted", nil)
}
