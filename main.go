package main

import (
	addmusicservice "github.com/Rammurthy5/music_store_gin_service/add_music_service"
	fetchmusicservice "github.com/Rammurthy5/music_store_gin_service/fetch_music_service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", fetchmusicservice.GetAlbums)
	router.POST("/albums", addmusicservice.PostAlbums)
	router.GET("/albums/:id", fetchmusicservice.GetAlbumById)
	router.Run("localhost:8080")
}
