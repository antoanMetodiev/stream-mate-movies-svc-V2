package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieDetails(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"data": "The Witcher: Season 4 + VIDEO URL",
	})
}
