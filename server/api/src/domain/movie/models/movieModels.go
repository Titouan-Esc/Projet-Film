package models

type ModelMovie struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Picture  string `json:"picture"`
}

type ModelMovieRequest struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Picture  string `json:"picture"`
}
