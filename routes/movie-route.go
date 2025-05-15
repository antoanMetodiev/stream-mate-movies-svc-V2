package routes

import (
	"movies-svc-V2/services"

	"github.com/gin-gonic/gin"
)

func MovieRoute(server *gin.Engine) {
	server.GET("/movie/The-Witcher", services.GetMovieDetails)
}
