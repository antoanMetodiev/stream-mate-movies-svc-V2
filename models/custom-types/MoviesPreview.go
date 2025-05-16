package customtypes

type MoviePreview struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	PosterImgUrl string `json:"posterImgURL"`
	ReleaseDate  string `json:"releaseDate"`
}
