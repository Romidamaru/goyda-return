package dto

import (
	"simple-api/internal/pkg/users/ent"
	"simple-api/internal/utils"
)

type CreateUser struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (dto *CreateUser) ToEntity() ent.User {
	return ent.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: utils.HashPassword(dto.Password),
	}
}
