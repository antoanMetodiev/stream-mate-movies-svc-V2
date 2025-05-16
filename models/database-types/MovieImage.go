package databasetypes

import (
	"github.com/google/uuid"
)

type MovieImage struct {
	ID        uuid.UUID `json:"id"`
	ImageType string    `json:"imageType"`
	ImageURL  string    `json:"imageURL"`
	MovieID   uuid.UUID `json:"movieId"`
}
