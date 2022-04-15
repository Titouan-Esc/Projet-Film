package movie

import (
	"fmt"
	"web-service/api/database"
	"web-service/api/src/domain/movie/models"
)

type MovieRepository struct {
	database.IDBHandler
}

func (entity *MovieRepository) SaveMovieDB(movie models.ModelMovie) error {
	return nil
}

func (entity *MovieRepository) GetAllMoviesDB() (models.ModelMovie, error) {

	var movie models.ModelMovie

	row, err := entity.Query(fmt.Sprintf(`SELECT "title", "overview", "picture" FROM "movies" ORDER BY "id" ASC`))
	if err != nil {
		return movie, err
	}

	for row.Next() {
		if err := row.Scan(&movie.Title, &movie.Overview, &movie.Picture); err != nil {
			return movie, err
		}
	}

	return movie, nil
}

func (entity *MovieRepository) GetMovieByIdDB(mid int) (models.ModelMovie, error) {
	var movie models.ModelMovie

	row, err := entity.Query(fmt.Sprintf(`SELECT "title", "overview", "picture" FROM "movies" WHERE "id" = %d`, mid))
	if err != nil {
		return movie, err
	}

	for row.Next() {
		if err := row.Scan(&movie.Title, &movie.Overview, &movie.Picture); err != nil {
			return movie, err
		}
	}

	return movie, nil
}

func (entity *MovieRepository) UpdateMovieDB(movie models.ModelMovie) (int, error) {
	var id int
	return id, nil
}

func (entity *MovieRepository) DeleteMovieDB(mid int) (int, error) {
	var id int
	return id, nil
}

func (entity *MovieRepository) MovieExistDB(title string) bool {
	var ouai bool
	return ouai
}
