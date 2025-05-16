package main

import (
	"log"
	"time"

	dbconnection "movies-svc-V2/db-connection"
	"movies-svc-V2/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	_, err := dbconnection.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	routes.ModularRoute(server)
	server.Run() // PORT 8080
}
