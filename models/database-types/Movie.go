package databasetypes

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title            string    `gorm:"not null" json:"title"`
	PosterImgURL     string    `gorm:"column:poster_img_url;not null" json:"posterImgURL"`
	SpecialText      string    `gorm:"column:special_text" json:"specialText"`
	SearchTag        string    `gorm:"column:search_tag;not null" json:"search_tag"`
	Genres           string    `json:"genres"`
	Description      string    `gorm:"type:text" json:"description"`
	ReleaseDate      string    `gorm:"column:release_date" json:"releaseDate"`
	TMDBRating       string    `gorm:"column:tmdb_rating" json:"tmdbRating"`
	BackgroundImgURL string    `gorm:"column:background_img_url" json:"backgroundImg_URL"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	VideoURL string `gorm:"column:video_url;not null" json:"videoURL"`

	CastList []Actor      `gorm:"many2many:movies_actors;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"castList"`
	Images   []MovieImage `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imagesList"`
}

func (MovieImage) TableName() string {
	return "movies_images"
}

func (Actor) TableName() string {
	return "actors" // ако името съвпада, може и да не е нужно
}

func (Movie) TableName() string {
	return "movies" // ако името съвпада, може и да не е нужно
}
