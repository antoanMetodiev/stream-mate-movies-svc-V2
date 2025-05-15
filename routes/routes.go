package routes

import "github.com/gin-gonic/gin"

func ModularRoute(server *gin.Engine) {
	// Routes:
	MovieRoute(server)
}
