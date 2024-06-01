package handler

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"
	"base-trade-rest/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

type ProductHandler struct {
	ProductService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	var productHandler = ProductHandler{ProductService: productService}
	return &productHandler
}

func (h *ProductHandler) Index(ctx *gin.Context) {

}

func (h *ProductHandler) Store(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt5.MapClaims)
	userID := uint(userData["id"].(float64))

	var request request.ProductCreateRequest

	err := ctx.ShouldBind(&request)
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

	product := model.Product{
		Name:     request.Name,
		ImageURL: nil,
		UserID:   userID,
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
