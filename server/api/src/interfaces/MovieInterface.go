package interfaces

import "web-service/api/src/domain/movie/models"

type IMovieRepository interface {
	// -----------------
	RecoverMoviesDB() ([]*models.ResultsModel, error)
	// -----------------
	LikeMovieDB(mid, like int, column string) (int, error)
	RecoverLikesMovieDB(mid int, column string) (int, error)
	// -----------------
	AddCommentMovieDB(mid int, comment string) (string, error)
	DeleteCommentMovieDB(mid int) (int, error)
}
