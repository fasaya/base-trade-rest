package middleware

import (
	"base-trade-rest/core/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			helpers.CreateFailedResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
