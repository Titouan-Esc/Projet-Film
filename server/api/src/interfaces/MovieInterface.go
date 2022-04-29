package interfaces

import "web-service/api/src/domain/movie/models"

type IMovieRepository interface {
	MovieExist(mid int) (bool, error)
	SaveMovieDB(movie models.ResultsModel) (int, error)
}
