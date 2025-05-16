package databasetypes

import (
	"github.com/google/uuid"
)

type Actor struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ImageURL          string    `gorm:"column:image_url" json:"imageURL"`
	Biography         string    `gorm:"type:text" json:"biography"`
	FacebookUsername  string    `gorm:"column:facebook_username" json:"facebookUsername"`
	InstagramUsername string    `gorm:"column:instagram_username" json:"instagramUsername"`
	TwitterUsername   string    `gorm:"column:twitter_username" json:"twitterUsername"`
	YoutubeChannel    string    `gorm:"column:youtube_channel" json:"youtubeChannel"`
	IMDBId            string    `gorm:"column:imdb_id" json:"imdbId"`
	Birthday          string    `json:"birthday"`
	KnownFor          string    `gorm:"column:known_for" json:"knownFor"`
	PlaceOfBirth      string    `gorm:"column:place_of_birth" json:"placeOfBirth"`
	Gender            string    `json:"gender"`
	Popularity        string    `json:"popularity"`
	NameInRealLife    string    `gorm:"column:name_in_real_life;not null" json:"nameInRealLife"`

	// Обратно ManyToMany към Movie
	Movies []Movie `gorm:"many2many:movies_actors;" json:"-"`
}
