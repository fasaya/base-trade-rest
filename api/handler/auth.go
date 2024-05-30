package handler

import (
	"base-trade-rest/api/request"
	"base-trade-rest/api/resource"
	"base-trade-rest/core/helpers"
	"base-trade-rest/core/model"
	"base-trade-rest/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AuthHandler struct {
	userService service.IUserService
}

func NewAuthHandler(userService service.IUserService) *AuthHandler {
	var authHandler = AuthHandler{userService: userService}
	return &authHandler
}

func (h *AuthHandler) UserRegister(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.CreateFailResponse(err.Error(), http.StatusBadRequest),
		)
		return
	}

	newUser, err := h.userService.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.CreateFailResponse(err.Error(), http.StatusBadRequest),
		)
		return
	}

	// Transform
	var authResource resource.AuthResource
	copier.Copy(&authResource, &newUser)

	ctx.JSON(
		http.StatusOK,
		helpers.CreateSuccessResponse("Registered successfully", http.StatusOK, authResource),
	)

}

func (h *AuthHandler) UserLogin(ctx *gin.Context) {
	var request request.AuthLoginRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.CreateFailResponse(err.Error(), http.StatusBadRequest),
		)
		return
	}

	user, err := h.userService.ValidateCredentials(request)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.CreateFailResponse(err.Error(), http.StatusBadRequest),
		)
		return
	}

	// Transform
	var authResource resource.AuthResource
	copier.Copy(&authResource, &user)

	authResource.Token = helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(
		http.StatusOK,
		helpers.CreateSuccessResponse("Login successfully", http.StatusOK, authResource),
	)
}
