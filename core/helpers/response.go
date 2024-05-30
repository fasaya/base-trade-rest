package helpers

type Response struct {
	Message    string              `json:"message"`
	Error      bool                `json:"error"`
	StatusCode uint                `json:"status_code"`
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

func CreateFailResponse(msg string, statusCode uint) Response {
	return Response{
		Error:      true,
		Message:    msg,
		Data:       nil,
		StatusCode: statusCode,
	}
}

func CreateSuccessResponse(msg string, statusCode uint, data any) Response {
	return Response{
		Error:      false,
		Message:    msg,
		Data:       data,
		StatusCode: statusCode,
	}
}

func CreatePaginatedResponse(msg string, statusCode uint, data any, pageMeta PaginationResponse) Response {
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
