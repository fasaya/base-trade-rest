package helpers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message    string              `json:"message"`
	Error      bool                `json:"error"`
	StatusCode int                 `json:"status_code"`
	Data       any                 `json:"data,omitempty"`
	Meta       *PaginationResponse `json:"meta,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type PaginationResponse struct {
	Page     int64 `json:"page"`
	Limit    int64 `json:"limit"`
	LastPage int64 `json:"last_page"`
	Total    int64 `json:"total"`
}

func CreateFailedResponse(ctx *gin.Context, statusCode int, msg string) {
	var res = Response{
		Error:      true,
		Message:    msg,
		Data:       nil,
		StatusCode: statusCode,
	}

	ctx.AbortWithStatusJSON(statusCode, res)
}

func CreateSuccessfulResponse(ctx *gin.Context, statusCode int, msg string, data any) {
	var res = Response{
		Error:      false,
		Message:    msg,
		Data:       data,
		StatusCode: statusCode,
	}

	ctx.JSON(statusCode, res)
}

func CreatePaginatedResponse(statusCode int, msg string, data any, pageMeta PaginationResponse) Response {
	return Response{
		Error:   false,
		Message: msg,
		Data:    data, Meta: &pageMeta,
		StatusCode: statusCode,
	}
}

func CreateAuthResponse(token string, role string) AuthResponse {
	return AuthResponse{
		Token: token,
		Role:  role,
	}
}
