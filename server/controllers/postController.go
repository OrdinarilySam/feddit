package controllers

import (
	"github.com/OrdinarilySam/feddit/initializers"
	"github.com/OrdinarilySam/feddit/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.Where("id = ?", c.Param("id")).First(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch post"})
		return
	}

	c.JSON(200, post)
}

func CreatePost(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var body struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	newPost := models.Post{
		Title:   body.Title,
		Content: body.Content,
		UserID:  user.ID,
	}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post created successfully"})
}

func UpdatePost(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var post models.Post
	result := initializers.DB.Where("id = ?", c.Param("id")).First(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch post"})
		return
	}

	if post.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var body struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	post.Title = body.Title
	post.Content = body.Content

	result = initializers.DB.Save(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var post models.Post
	result := initializers.DB.Where("id = ?", c.Param("id")).First(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch post"})
		return
	}

	if post.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

func DeleteAllPosts(c *gin.Context) {
	result := initializers.DB.Exec("DELETE FROM posts")
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete posts"})
		return
	}

	c.JSON(200, gin.H{"message": "All posts deleted successfully"})
}
