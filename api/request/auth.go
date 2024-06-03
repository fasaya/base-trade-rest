package request

type AuthLoginRequest struct {
	Email    string `json:"email" form:"email" valid:"required,email"`
	Password string `json:"password" form:"password" valid:"required"`
}

type AuthRegisterRequest struct {
	Name     string `json:"name" form:"name" valid:"required"`
	Email    string `json:"email" form:"email" valid:"required,email"`
	Password string `json:"password" form:"password" valid:"required"`
}
