package domain

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string
	Content string
}
