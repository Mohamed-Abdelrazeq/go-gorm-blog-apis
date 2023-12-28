package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string
	Content string
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Blog{})
	return db
}

func main() {
	db := initDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hello from %s", hostname)
	})

	// get all blogs
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		var blogs []Blog
		db.Find(&blogs)
		fmt.Fprintf(w, "Blogs: %v", blogs)
	})

	// create a blog
	http.HandleFunc("/blog/create", func(w http.ResponseWriter, r *http.Request) {
		db.Create(&Blog{Title: "Hello", Content: "World"})
		fmt.Fprintf(w, "Blog created")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
