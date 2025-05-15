package main

import (
	"log"

	dbconnection "movies-svc-V2/db-connection"
	"movies-svc-V2/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	_, err := dbconnection.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	routes.ModularRoute(server)
	server.Run() // PORT 8080
}
