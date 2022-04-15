package movie

import (
	"net/http"
	"web-service/api/src/interfaces"
)

type MovieController struct {
	interfaces.IMovieRepository
}

func (controller *MovieController) GetAllMovies(res http.ResponseWriter, req *http.Request) {}

func (controller *MovieController) GetOneMovie(res http.ResponseWriter, req *http.Response) {}

func (controller *MovieController) CreateMovie(res http.ResponseWriter, req *http.Request) {}

func (controller *MovieController) UpdateMovie(res http.ResponseWriter, req *http.Request) {}

func (controller *MovieController) DeleteMovie(res http.ResponseWriter, req *http.Request) {}
