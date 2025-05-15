package services

import (
	dbconnection "movies-svc-V2/db-connection"
	databasetypes "movies-svc-V2/models/database-types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieDetails(c *gin.Context) {
	movieTitle := "Anora" // или можеш да вземеш от параметър: c.Param("title")

	var movie databasetypes.Movie

	err := dbconnection.DB.
		Preload("CastList").
		Preload("Images").
		Where("title = ?", movieTitle).First(&movie).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}
