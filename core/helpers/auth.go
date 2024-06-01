package helpers

import (
	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

type AuthUser struct {
	ID    uint
	Email string
}

func GetAuthUser(ctx *gin.Context) AuthUser {

	userData := ctx.MustGet("userData").(jwt5.MapClaims)

	return AuthUser{
		ID:    uint(userData["id"].(float64)),
		Email: userData["email"].(string),
	}
}
