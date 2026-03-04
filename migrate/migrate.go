package main

import (
	initializers "github.com/GoProject/go-crud/Initializers"
	models "github.com/GoProject/go-crud/Models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
}
