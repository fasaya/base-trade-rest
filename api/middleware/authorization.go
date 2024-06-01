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
		productUUID := ctx.Param("uuid")

		userData := ctx.MustGet("userData").(jwt5.MapClaims)
		userID := uint(userData["id"].(float64))

		var findProduct = []request.FieldValueRequest{
			{Field: "uuid", Value: productUUID},
			{Field: "user_id", Value: userID},
		}

		productRepository := repository.NewProductRepository(db)

		product, err := productRepository.GetProductByMultipleKey(findProduct)
		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusNotFound, "Data Not Found")
			return
		}

		if product.UserID != userID {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, "You are not allowed to access this data")
			return
		}

		ctx.Next()
	}
}
