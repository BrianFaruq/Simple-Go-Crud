package main

import (
	initializers "github.com/GoProject/go-crud/Initializers"
	"github.com/GoProject/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.CreatePosts)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.ShowPost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.LogIn)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
