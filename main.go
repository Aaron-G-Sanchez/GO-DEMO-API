package main

import (
	"github.com/aaron-g-sanchez/go-crud/controllers"
	"github.com/aaron-g-sanchez/go-crud/initializers"
	"github.com/aaron-g-sanchez/go-crud/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/albums", controllers.FindAllAlbums)

	router.POST("/albums", controllers.CreateAlbum)

	router.Run()
}
