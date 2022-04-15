package interfaces

import "web-service/api/src/domain/movie/models"

type IMovieRepository interface {
	SaveMovieDB(movie models.ModelMovie) error
	GetAllMoviesDB() (models.ModelMovie, error)
	GetMovieByIdDB(mid int) (models.ModelMovie, error)
	UpdateMovieDB(movie models.ModelMovieRequest) (int, error)
	DeleteMovieDB(mid int) (int, error)
	MovieExistDB(title string) bool
}
