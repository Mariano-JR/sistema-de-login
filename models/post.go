package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Titulo     string `json:"titulo"`
	Conteudo   string `json:"conteudo"`
	HoraDoPost string `json:"horadopost"`
}
