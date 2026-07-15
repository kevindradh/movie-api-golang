package handlers

import (
	"MovieAPI/config"
	"MovieAPI/models"
	"MovieAPI/response"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetMovies godoc
// @Summary 		Retrieve all movies
// @Description 	Retrieve list of movies, can filter by genre or title
// @Tags 			movies
// @Accept 			json
// @Produce 		json
// @Param 			genre query string false "Filter by genre"
// @Param 			title query string false "Filter by title"
// @Success			200	{object} response.ListResponse
// @Failure			500 {object} response.ErrorResponse
// @Router 			/movies [get]
func GetMovies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}

	genre := c.Query("genre")
	title := c.Query("title")

	if genre != "" {
		filter["genre"] = genre
	}
	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"}
	}

	cursor, err := config.MovieCollection.Find(ctx, filter)
	if err != nil {
		response.InternalError(c, "Failed to fetch movies", err)
		return
	}
	defer cursor.Close(ctx)

	var movies []models.Movie
	if err := cursor.All(ctx, &movies); err != nil {
		response.InternalError(c, "Failed to fetch movies", err)
		return
	}

	response.List(c, "Success retreive movies", movies)
}

// GetMovieByID godoc
// @Summary 		Retrieve detail movie
// @Description 	Retrieve one movie details by movie ID
// @Tags 			movies
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Movie ID"
// @Success 		200 {object} response.SuccessResponse
// @Failure			400 {object} response.ErrorResponse
// @Failure			404 {object} response.ErrorResponse
// @Router 			/movies/{id} [get]
func GetMovieByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err)
		return
	}

	var movie models.Movie
	err = config.MovieCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie)
	if err != nil {
		response.NotFound(c, "Movie not found")
		return
	}

	response.OK(c, "Success retrieve movie", movie)
}

// CreateMovie godoc
// @Summary 		Insert movie
// @Description 	Insert a new movie
// @Tags 			movies
// @Accept 			json
// @Produce 		json
// @Param 			movie body models.Movie true "New movie"
// @Success			201	{object} response.SuccessResponse
// @Failure			400 {object} response.ErrorResponse
// @Failure			500 {object} response.ErrorResponse
// @Router 			/movies [post]
func CreateMovie(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var newMovie models.Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		response.BadRequest(c, "Invalid request", err)
		return
	}

	result, err := config.MovieCollection.InsertOne(ctx, newMovie)
	if err != nil {
		response.InternalError(c, "Failed to create movie", err)
		return
	}

	newMovie.ID = result.InsertedID.(primitive.ObjectID)
	response.Created(c, "Success create movie", newMovie)
}

// UpdateMovie godoc
// @Tags 			movies
// @Summary 		Update movie
// @Description 	Update movie by movie ID
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Movie ID"
// @Param 			movie body models.Movie true "Updated movie"
// @Success 		200 {object} response.SuccessResponse
// @Failure 		400 {object} response.ErrorResponse
// @Failure 		404 {object} response.ErrorResponse
// @Router 			/movies/{id} [put]
func UpdateMovie(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err)
		return
	}

	var updateMovie models.Movie
	if err := c.ShouldBindJSON(&updateMovie); err != nil {
		response.BadRequest(c, "Invalid request", err)
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":    updateMovie.Title,
			"director": updateMovie.Director,
			"year":     updateMovie.Year,
			"genre":    updateMovie.Genre,
			"rating":   updateMovie.Rating,
			"poster":   updateMovie.Poster,
		},
	}

	result, err := config.MovieCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		response.InternalError(c, "Failed to update movie", err)
		return
	}

	if result.ModifiedCount == 0 {
		response.NotFound(c, "Movie not found")
		return
	}

	updateMovie.ID = objID
	response.OK(c, "Success update movie", updateMovie)
}

// DeleteMovie godoc
// @Summary 		Delete movie
// @Description 	Delete movie by movie ID
// @Tags 			movies
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Movie ID"
// @Success 		200 {object} response.SuccessResponse
// @Failure 		400 {object} response.ErrorResponse
// @Failure 		404 {object} response.ErrorResponse
// @Router 			/movies/{id} [delete]
func DeleteMovie(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err)
		return
	}

	result, err := config.MovieCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		response.InternalError(c, "Failed to delete movie", err)
		return
	}

	if result.DeletedCount == 0 {
		response.NotFound(c, "Movie not found")
		return
	}

	response.OK(c, "Success delete movie", nil)
}
