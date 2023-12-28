package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := db.InitDB()

	// create a gin router with default middleware
	router := gin.Default()

	// test index route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	// get all blogs
	// router.GET("/blogs", func(c *gin.Context) {
	// 	var blogs []Blog
	// 	db.Find(&blogs)
	// 	c.JSON(http.StatusOK, gin.H{"data": blogs})
	// })

	// create a blog
	// router.POST("/blogs", func(c *gin.Context) {
	// 	var input Blog
	// 	if err := c.ShouldBindJSON(&input); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	blog := Blog{Title: input.Title, Content: input.Content}
	// 	db.Create(&blog)
	// 	c.JSON(http.StatusOK, gin.H{"data": blog})
	// })

	// serve the application on port 3000
	log.Fatal(router.Run(":3000"))
}
