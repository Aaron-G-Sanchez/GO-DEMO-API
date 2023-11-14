package main

import (
	"github.com/aaron-g-sanchez/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{
		ID:     "1",
		Title:  "Torches",
		Artist: "Foster the People",
		Price:  10.99,
	},
	{
		ID:     "2",
		Title:  "SMITHEREENS",
		Artist: "Joji",
		Price:  24.99,
	},
	{
		ID:     "3",
		Title:  "Dark Sun",
		Artist: "Day Seeker",
		Price:  15.99,
	},
}

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.POST("/albums", postAlbums)

	router.Run()
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(200, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(200, newAlbum)
}
