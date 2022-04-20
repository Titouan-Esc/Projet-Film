package models

type MovieModel struct {
	Genres []GenresModel `json:"genres"`
}

type GenresModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
