package app

import (
	domain "mohamed-abdelrazeq/blogger/pkg/domain"
)

type BlogApp struct {
	db domain.BlogDB
}

// NewBlogApp creates a new instance of BlogApp
func NewBlogApp(db domain.BlogDB) domain.BlogApp {
	return &BlogApp{db}
}

// Get retrieves a blog by ID
func (app *BlogApp) Get(id uint) (domain.Blog, error) {
	return app.db.Get(id)
}

// List retrieves all blogs
func (app *BlogApp) List() ([]domain.Blog, error) {
	return app.db.List()
}

// Create creates a new blog
func (app *BlogApp) Create(blog domain.Blog) (domain.Blog, error) {
	return app.db.Create(blog)
}

// Update updates a blog by ID
func (app *BlogApp) Update(id uint, blog domain.Blog) (domain.Blog, error) {
	return app.db.Update(id, blog)
}

// Delete deletes a blog by ID
func (app *BlogApp) Delete(id uint) (domain.Blog, error) {
	return app.db.Delete(id)
}
