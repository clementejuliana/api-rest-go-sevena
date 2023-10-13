package models

import "gorm.io/gorm"

type Estado struct {
	gorm.Model
	Status   string `json:"status,omitempty"`
	Nome string `json:"nome,omitempty"`
}