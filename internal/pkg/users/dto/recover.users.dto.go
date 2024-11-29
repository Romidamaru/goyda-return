package dto

type RecoverPasswordUser struct {
	Password    string `json:"password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}
