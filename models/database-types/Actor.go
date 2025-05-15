package databasetypes

import (
	"github.com/google/uuid"
)

type Actor struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ImageURL          string    `gorm:"column:image_url" json:"image_url"`
	Biography         string    `gorm:"type:text" json:"biography"`
	FacebookUsername  string    `gorm:"column:facebook_username" json:"facebook_username"`
	InstagramUsername string    `gorm:"column:instagram_username" json:"instagram_username"`
	TwitterUsername   string    `gorm:"column:twitter_username" json:"twitter_username"`
	YoutubeChannel    string    `gorm:"column:youtube_channel" json:"youtube_channel"`
	IMDBId            string    `gorm:"column:imdb_id" json:"imdb_id"`
	Birthday          string    `json:"birthday"`
	KnownFor          string    `gorm:"column:known_for" json:"known_for"`
	PlaceOfBirth      string    `gorm:"column:place_of_birth" json:"place_of_birth"`
	Gender            string    `json:"gender"`
	Popularity        string    `json:"popularity"`
	NameInRealLife    string    `gorm:"column:name_in_real_life;not null" json:"name_in_real_life"`

	// Обратно ManyToMany към Movie
	Movies []Movie `gorm:"many2many:movies_actors;" json:"-"`
}
