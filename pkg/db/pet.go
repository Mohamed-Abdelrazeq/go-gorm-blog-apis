package db

import (
	"mohamed-abdelrazeq/blogger/pkg/domain"

	"gorm.io/gorm"
)

type PetDB struct {
	db *gorm.DB
}

// NewPetDB creates a new instance of PetDB
func NewPetDB(db *gorm.DB) domain.BlogDB {
	return &PetDB{db}
}

// Get retrieves a blog by ID
func (db *PetDB) Get(id uint) (domain.Blog, error) {
	var blog domain.Blog
	result := db.db.First(&blog, id)
	return blog, result.Error
}

// List retrieves all blogs
func (db *PetDB) List() ([]domain.Blog, error) {
	var blogs []domain.Blog
	result := db.db.Find(&blogs)
	return blogs, result.Error
}

// Create creates a new blog
func (db *PetDB) Create(blog domain.Blog) (domain.Blog, error) {
	result := db.db.Create(&blog)
	return blog, result.Error
}

// Update updates a blog by ID
func (db *PetDB) Update(id uint, blog domain.Blog) (domain.Blog, error) {
	result := db.db.Model(&blog).Where("id = ?", id).Updates(blog)
	return blog, result.Error
}

// Delete deletes a blog by ID
func (db *PetDB) Delete(id uint) (domain.Blog, error) {
	var blog domain.Blog
	result := db.db.Delete(&blog, id)
	return blog, result.Error
}
