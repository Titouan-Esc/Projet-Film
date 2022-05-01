package interfaces

import "web-service/api/src/domain/movie/models"

type IMovieRepository interface {
	MovieExist(mid int) (bool, error)
	SaveMovieDB(movie models.ResultsModel, overview, title string) (models.ResultsModel, error)
	LikeMovieDB(mid, like int, column string) (int, error)
	RecoverLikesMovieDB(mid int, column string) (int, error)
	AddCommentMovieDB(mid int, txt string) (string, error)
	UpdateCommentMovieDB(mid int, txt string) (string, error)
	DeleteCommentMovieDB(mid int) (int, error)
}
