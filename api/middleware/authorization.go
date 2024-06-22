package middleware

import (
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/repository"
	"base-trade-rest/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productUUID := ctx.Param("productUUID")

		// Get authenticated user
		userData := helpers.GetAuthUser(ctx)

		productRepository := repository.NewProductRepository(db)

		product, err := productRepository.GetProductByKey("uuid", productUUID)
		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "Product not exist")
			return
		}

		if product.UserID != userData.ID {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx.Next()
	}
}

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		variantUUID := ctx.Param("variantUUID")

		// Get authenticated user
		userData := helpers.GetAuthUser(ctx)

		variantRepository := repository.NewVariantRepository(db)
		variant, err := variantRepository.GetVariantByKey("uuid", variantUUID)
		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "Variant not exist")
			return
		}

		if variant.Product.UserID != userData.ID {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx.Next()
	}
}
