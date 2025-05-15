package databasetypes

import (
	"github.com/google/uuid"
)

type MovieImage struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ImageType string    `gorm:"column:image_type;not null"`
	ImageURL  string    `gorm:"column:image_url;not null"`

	MovieID uuid.UUID `gorm:"column:movie_id;not null"` // Foreign key
}
