package addmusicservice

import (
	"net/http"

	datastore "github.com/Rammurthy5/music_catalog_service/data_store"
	"github.com/gin-gonic/gin"
)

func PostAlbums(c *gin.Context) {
	var newAlbum datastore.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	datastore.Albums = append(datastore.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
