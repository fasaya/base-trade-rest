package handler

import (
	"base-trade-rest/api/request"
	"base-trade-rest/api/resource"
	"base-trade-rest/core/helpers"
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
	var user request.AuthRegisterRequest

	err := ctx.ShouldBind(&user)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := h.userService.CreateUser(&user)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Transform
	var authResource resource.AuthResource
	copier.Copy(&authResource, &newUser)

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Registered successfully", authResource)
}

func (h *AuthHandler) UserLogin(ctx *gin.Context) {
	var request request.AuthLoginRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userService.ValidateCredentials(request)
	if err != nil {
		helpers.CreateFailedResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Transform
	var authResource resource.AuthResource
	copier.Copy(&authResource, &user)

	authResource.Token = helpers.GenerateToken(user.ID, user.Email)

	helpers.CreateSuccessfulResponse(ctx, http.StatusOK, "Login successfully", authResource)
}
