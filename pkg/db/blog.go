package db

import (
	"mohamed-abdelrazeq/blogger/pkg/domain"

	"gorm.io/gorm"
)

type BlogDB struct {
	db *gorm.DB
}

// NewBlogDB creates a new instance of BlogDB
func NewBlogDB(db *gorm.DB) domain.BlogDB {
	return &BlogDB{db}
}

// Get retrieves a blog by ID
func (db *BlogDB) Get(id uint) (domain.Blog, error) {
	var blog domain.Blog
	result := db.db.First(&blog, id)
	return blog, result.Error
}

// List retrieves all blogs
func (db *BlogDB) List() ([]domain.Blog, error) {
	var blogs []domain.Blog
	result := db.db.Find(&blogs)
	return blogs, result.Error
}

// Create creates a new blog
func (db *BlogDB) Create(blog domain.Blog) (domain.Blog, error) {
	result := db.db.Create(&blog)
	return blog, result.Error
}

// Update updates a blog by ID
func (db *BlogDB) Update(id uint, blog domain.Blog) (domain.Blog, error) {
	result := db.db.Model(&blog).Where("id = ?", id).Updates(blog)
	return blog, result.Error
}

// Delete deletes a blog by ID
func (db *BlogDB) Delete(id uint) (domain.Blog, error) {
	var blog domain.Blog
	result := db.db.Delete(&blog, id)
	return blog, result.Error
}
