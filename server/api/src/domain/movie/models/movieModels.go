package models

// type MovieModel struct {
// 	Genres []GenresModel `json:"genres"`
// }

// type GenresModel struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

type DiscoverModel struct {
	Page         int            `json:"page"`
	Results      []ResultsModel `json:"results"`
	TotalResults int            `json:"total_results"`
	TotalPages   int            `json:"total_pages"`
}

type ResultsModel struct {
	ID          int     `json:"id"`
	PosterPath  string  `json:"poster_path"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	Title       string  `json:"title"`
	Popularity  float32 `json:"popularity"`
}
