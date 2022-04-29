package movie

import (
	"web-service/api/database"
	"web-service/api/src/domain/movie/models"
)

type MovieRepository struct {
	database.IDBHandler
}

func (repository *MovieRepository) MovieExist(mid int) (bool, error) {
	var ouai bool
	return ouai, nil
}

func (repository *MovieRepository) SaveMovieDB(movie models.ResultsModel) (int, error) {
	var id int
	return id, nil
}
