package fetchmusicservice

import (
	"net/http"

	datastore "github.com/Rammurthy5/music_catalog_service/data_store"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datastore.Albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range datastore.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
