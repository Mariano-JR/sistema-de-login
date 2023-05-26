package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"max=15, nonzero"`
	Email    string `json:"email" validate:"nonzero"`
	Senha    string `json:"senha" validate:"max=15, nonzero"`
}
