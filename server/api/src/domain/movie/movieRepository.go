package movie

import (
	"fmt"
	"web-service/api/database"
	"web-service/api/src/domain/movie/models"
)

type MovieRepository struct {
	database.IDBHandler
}

func (repository *MovieRepository) RecoverMoviesDB() ([]*models.ResultsModel, error) {
	movies := make([]*models.ResultsModel, 0)
	row, err := repository.Query(fmt.Sprintf(`SELECT * FROM "movies" ORDER BY "id" ASC`))
	if err != nil {
		return movies, err
	}

	for row.Next() {
		movie := new(models.ResultsModel)
		if err := row.Scan(&movie.ID, &movie.PosterPath, &movie.Overview, &movie.ReleaseDate, &movie.Title, &movie.Popularity, &movie.Likes, &movie.Dislikes, &movie.Comments); err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (repository *MovieRepository) LikeMovieDB(mid, like int, column string) (int, error) {
	row, err := repository.Query(fmt.Sprintf(`UPDATE "movies" SET "%s" = %d WHERE "id" = %d RETURNING "%s"`, column, like, mid, column))
	if err != nil {
		return -1, err
	}

	var likes int
	for row.Next() {
		if err := row.Scan(&likes); err != nil {
			return -1, err
		}
	}

	return likes, nil
}

func (reposiroty *MovieRepository) RecoverLikesMovieDB(mid int, column string) (int, error) {
	row, err := reposiroty.Query(fmt.Sprintf(`SELECT "%s" FROM "movies" WHERE "id" = %d`, column, mid))
	if err != nil {
		return -1, err
	}

	var likes int
	for row.Next() {
		if err := row.Scan(&likes); err != nil {
			return -1, err
		}
	}

	return likes, nil
}

func (repository *MovieRepository) AddCommentMovieDB(mid int, comment string) (string, error) {
	var commentaire string

	row, err := repository.Query(fmt.Sprintf(`UPDATE "movies" SET "comments" = '%s' WHERE "id" = %d RETURNING "comments"`, comment, mid))
	if err != nil {
		return commentaire, err
	}

	for row.Next() {
		if err := row.Scan(&commentaire); err != nil {
			return commentaire, err
		}
	}

	return commentaire, nil
}

func (repository *MovieRepository) DeleteCommentMovieDB(mid int) (int, error) {
	row, err := repository.Query(fmt.Sprintf(`UPDATE "movies" SET "comments" = '' WHERE "id" = %d RETURNING "id"`, mid))
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
