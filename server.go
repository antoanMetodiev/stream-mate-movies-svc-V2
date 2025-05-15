package main

import (
	"movies-svc-V2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.ModularRoute(server)
	server.Run() // PORT 8080
}
