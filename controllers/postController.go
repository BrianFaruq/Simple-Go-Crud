package controllers

import (
	initializers "github.com/GoProject/go-crud/Initializers"
	models "github.com/GoProject/go-crud/Models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {

	//Body struct

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Creating post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func ShowPost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	//return the posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	//Body struct
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	//Delete
	initializers.DB.Delete(&models.Post{}, id)

	//return
	c.Status(200)
}
