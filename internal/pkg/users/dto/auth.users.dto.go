package dto

type LoginUser struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}
