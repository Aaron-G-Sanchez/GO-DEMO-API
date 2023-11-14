package controllers

import (
	"github.com/aaron-g-sanchez/go-crud/models"
	"github.com/gin-gonic/gin"
)

func FindAllAlbums(c *gin.Context) {
	var albums []models.Album

	models.DB.Find(&albums)

	c.IndentedJSON(200, gin.H{"data": albums})
}

type CreateAlbumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

func CreateAlbum(c *gin.Context) {
	var input CreateAlbumInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	models.DB.Create(&album)

	c.IndentedJSON(200, gin.H{"data": album})
}
