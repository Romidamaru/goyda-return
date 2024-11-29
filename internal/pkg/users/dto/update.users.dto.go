package dto

type UpdateUser struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
}
