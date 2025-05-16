package routes

import (
	"movies-svc-V2/services"

	"github.com/gin-gonic/gin"
)

func MovieRoute(server *gin.Engine) {

	// GET DATA (Just to consume data!)
	server.GET("/get-movie-details", services.GetMovieDetails)
	server.GET("/get-all-movies-count", services.GetAllMoviesCount)
	server.GET("/get-next-thirty-movies", services.GetNextThirtyMovies)
	server.GET("/get-movies-count-by-genre", services.GetAllMoviesCountByGenre)
	server.GET("/get-next-twenty-movies-by-genre", services.GetAllMoviesByGenre)
	server.GET("/get-movies-by-title", services.GetMoviesByTitle)
	server.GET("/get-searched-movies-count", services.GetMoviesCountByTitle)

	// CREATE DATA (Create records in database!)
}
