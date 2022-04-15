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

func (entity *MovieRepository) UpdateMovieDB(movie models.ModelMovieRequest) (int, error) {
	_, err := entity.Query(fmt.Sprintf(`UPDATE "movies" SET "title" = '%s', "overview" = '%s', "picture" = '%s' WHERE "id" = %d`, movie.Title, movie.Overview, movie.Picture, movie.ID))
	if err != nil {
		return -1, err
	}

	return movie.ID, nil
}

func (entity *MovieRepository) DeleteMovieDB(mid int) (int, error) {
	row, err := entity.Query(fmt.Sprintf(`DELETE FROM "movies" WHERE "id" + %d RETURNING "id"`, mid))
	if err != nil {
		return -1, err
	}

	var id int
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (entity *MovieRepository) MovieExistDB(title string) bool {
	row, err := entity.Query(fmt.Sprintf(`SELECT "id" FROM "movies" WHERE "title" = '%s'`, title))
	if err != nil {
		return true
	}

	var id int
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return true
		}
	}

	if id > 1 {
		return true
	}

	return false
}
