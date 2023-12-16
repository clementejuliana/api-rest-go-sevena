package models

import "gorm.io/gorm"

type TipoAtividade struct {
    gorm.Model
    ID int `json:"id"`
    Nome string `json:"nome"`
}