package services

import (
	"fmt"
	dbconnection "movies-svc-V2/db-connection"
	"net/http"
	"strconv"

	customtypes "movies-svc-V2/models/custom-types"
	databasetypes "movies-svc-V2/models/database-types"

	"github.com/gin-gonic/gin"
)

func GetMoviesCountByTitle(context *gin.Context) {
	title := context.Query("movieName")
	likePattern := fmt.Sprintf("%%%s%%", title)

	var count int64
	err := dbconnection.DB.Raw(`
		SELECT count(*) FROM movies 
		WHERE LOWER(title) LIKE LOWER(?) 
		   OR LOWER(search_tag) LIKE LOWER(?)`,
		likePattern, likePattern).
		Scan(&count).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting movies count"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"count": count})
}

func GetMoviesByTitle(context *gin.Context) {
	title := context.Query("title")
	likePattern := fmt.Sprintf("%%%s%%", title)
	fmt.Println(title + "=>>>>>>>>>>>>>>>>>>")

	var movies []databasetypes.Movie

	err := dbconnection.DB.Raw(`
		SELECT * FROM movies 
		WHERE LOWER(title) LIKE LOWER(?) 
		   OR LOWER(search_tag) LIKE LOWER(?)`,
		likePattern, likePattern).
		Scan(&movies).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching movies"})
		return
	}

	context.JSON(http.StatusOK, movies)
}

func GetAllMoviesByGenre(context *gin.Context) {
	genre := context.Query("receivedGenre")
	pageStr := context.DefaultQuery("page", "0")
	sizeStr := context.DefaultQuery("size", "20")

	likePattern := fmt.Sprintf("%%%s%%", genre)

	// Парсваме към int
	page, err1 := strconv.Atoi(pageStr)
	size, err2 := strconv.Atoi(sizeStr)

	if err1 != nil || err2 != nil || page < 0 || size <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	offset := page * size

	var movies []databasetypes.Movie
	err := dbconnection.DB.
		Raw(`SELECT * FROM movies WHERE LOWER(genres) LIKE LOWER(?) LIMIT ? OFFSET ?`,
			likePattern, size, offset).
		Scan(&movies).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Movies genre not found!"})
		return
	}

	context.JSON(http.StatusOK, movies)
}

func GetAllMoviesCountByGenre(context *gin.Context) {
	genre := context.Query("receivedGenre")
	likePattern := fmt.Sprintf("%%%s%%", genre) // => "%Action%" например

	var count int64
	err := dbconnection.DB.
		Raw(`SELECT COUNT(*) FROM movies WHERE LOWER(genres) LIKE LOWER(?)`, likePattern).
		Scan(&count).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Movies genre not found!"})
		return
	}

	context.JSON(http.StatusOK, count)
}

func GetMovieDetails(context *gin.Context) {
	cinemaRecordId := context.Query("id") // или можеш да вземеш от параметър: c.Param("title")

	var movie databasetypes.Movie
	err := dbconnection.DB.
		Preload("CastList").
		Preload("Images").
		First(&movie, "id = ?", cinemaRecordId).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Movie not found"})
		return
	}

	context.JSON(http.StatusOK, movie)
}

func GetAllMoviesCount(context *gin.Context) {
	var movieCount = 0
	err := dbconnection.DB.Raw("SELECT COUNT(*) FROM movies").Scan(&movieCount).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}

	context.JSON(http.StatusOK, movieCount)
}

func GetNextThirtyMovies(context *gin.Context) {
	pageStr := context.DefaultQuery("page", "0")
	sizeStr := context.DefaultQuery("size", "20")

	page, err1 := strconv.Atoi(pageStr)
	size, err2 := strconv.Atoi(sizeStr)

	if err1 != nil || err2 != nil || page < 0 || size <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	offset := page * size
	var movies []customtypes.MoviePreview

	err := dbconnection.DB.
		Raw(`
			SELECT id, title, poster_img_url, release_date 
			FROM movies 
			ORDER BY created_at DESC 
			LIMIT ? OFFSET ?`, size, offset).
		Scan(&movies).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	context.JSON(http.StatusOK, movies)
}
