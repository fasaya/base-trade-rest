package middleware

import (
	"base-trade-rest/api/request"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/repository"
	"base-trade-rest/database"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productUUID := ctx.Param("productUUID")

		userData := ctx.MustGet("userData").(jwt5.MapClaims)
		userID := uint(userData["id"].(float64))

		var findProduct = []request.FieldValueRequest{
			{Field: "uuid", Value: productUUID},
			{Field: "user_id", Value: userID},
		}

		productRepository := repository.NewProductRepository(db)

		_, err := productRepository.GetProductByMultipleKey(findProduct)
		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "Data not exist or you're not authorized")
			return
		}

		ctx.Next()
	}
}

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productUUID := ctx.Param("productUUID")

		userData := ctx.MustGet("userData").(jwt5.MapClaims)
		userID := uint(userData["id"].(float64))

		var findVariant = []request.FieldValueRequest{
			{Field: "uuid", Value: productUUID},
			{Field: "user_id", Value: userID},
		}

		variantRepository := repository.NewVariantRepository(db)

		_, err := variantRepository.GetVariantByMultipleKey(findVariant)
		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "Data not exist or you're not authorized")
			return
		}

		ctx.Next()
	}
}
