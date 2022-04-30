package movie

import (
	"fmt"
	"web-service/api/database"
	"web-service/api/src/domain/movie/models"
)

type MovieRepository struct {
	database.IDBHandler
}

func (repository *MovieRepository) MovieExist(mid int) (bool, error) {
	row, err := repository.Query(fmt.Sprintf(`SELECT "id" FROM "movies" WHERE "id" = %d`, mid))
	if err != nil {
		return true, err
	}

	var id int
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return true, err
		}
	}

	if id > 1 {
		return true, nil
	}

	return false, nil
}

func (repository *MovieRepository) SaveMovieDB(movie models.ResultsModel, overview, title string) (models.ResultsModel, error) {
	var movies models.ResultsModel

	row, err := repository.Query(fmt.Sprintf(`INSERT INTO "movies" ("id", "poster_path", "overview", "release_date", "title", "popularity") VALUES (%d, '%s', '%s', '%s', '%s', %e) RETURNING *`, movie.ID, movie.PosterPath, overview, movie.ReleaseDate, title, movie.Popularity))
	if err != nil {
		return movies, err
	}

	for row.Next() {
		if err := row.Scan(&movies.ID, &movies.PosterPath, &movies.Overview, &movies.ReleaseDate, &movies.Title, &movies.Popularity); err != nil {
			return movies, err
		}
	}

	return movies, nil
}
