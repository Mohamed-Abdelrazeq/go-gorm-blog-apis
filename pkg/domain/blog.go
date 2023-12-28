package domain

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string
	Content string
}

type BlogApp interface {
	Get(id uint) (Blog, error)
	List() ([]Blog, error)
	Create(Blog) (Blog, error)
	Update(id uint, blog Blog) (Blog, error)
	Delete(id uint) (Blog, error)
}

type BlogDB interface {
	Get(id uint) (Blog, error)
	List() ([]Blog, error)
	Create(Blog) (Blog, error)
	Update(id uint, blog Blog) (Blog, error)
	Delete(id uint) (Blog, error)
}
