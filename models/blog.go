package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Body string `gorm:"type:text" json:"body" binding:"required"`
}